import os
import json
import time
import logging
from collections import defaultdict
from groq import Groq
import instructor
from dotenv import load_dotenv
from ..models.code_descriptor_model import CodeDescriptorModel, sys_prompt as code_descriptor_sys_prompt
from ..models.code_quality_eval_model import CodeQualityModel, sys_prompt as code_quality_sys_prompt
from ..models.code_sec_eval_model import CodeSecurityModel, sys_prompt as code_sec_sys_prompt
load_dotenv(dotenv_path=".env")
logging.basicConfig(level=logging.INFO, format='%(levelname)s - %(message)s')

GROQ_MODEL = "gemma2-9b-it"#"llama-3.1-70b-versatile"
TIME_DELAY = 1

class CodeAnalyser:
    def __init__(self, analysis_mode: str):
        self.logger = logging.getLogger(__name__)
        self.sys_prompt = None
        self.response_model = None
        if analysis_mode == "code_descriptor":
            self.sys_prompt = code_descriptor_sys_prompt
            self.response_model = CodeDescriptorModel
        elif analysis_mode == "code_quality":
            self.sys_prompt = code_quality_sys_prompt
            self.response_model = CodeQualityModel
        elif analysis_mode == "code_security":
            self.sys_prompt = code_sec_sys_prompt
            self.response_model = CodeSecurityModel
        else:
            print("choose one of the following analysis modes: code_descriptor, code_quality, code_security")

    def get_code(self, file_path: str):
        with open(file_path, 'r') as f:
            return f.read()

    def get_output(self, file_path: str):
        client = Groq(api_key=os.getenv("GROQ_API_KEY"))
        client = instructor.from_groq(client, mode=instructor.Mode.TOOLS)

        output = client.chat.completions.create(
            model=GROQ_MODEL,
            messages=[
                {"role": "system", "content": self.sys_prompt},
                {"role": "user", "content": self.get_code(file_path)},
            ],
            response_model=self.response_model,
        )
        return output

    def process_repo(self, repo_path):
        mapping = {}
        output_folder = os.path.join(repo_path, "output_data")
        os.makedirs(output_folder, exist_ok=True)
        chunk_folder_path = os.path.join(repo_path, "chunk_data")

        for file in os.listdir(chunk_folder_path):
            file_path = os.path.join(chunk_folder_path, file)
            if os.path.isfile(file_path):
                self.logger.info(f"Processing chunk: {file_path}")
                self.process_chunk(file_path, output_folder, mapping)
                time.sleep(TIME_DELAY)

        with open(os.path.join(repo_path, "file_output_mapping.json"), "w") as f:
            json.dump(mapping, f, indent=2)

        self.final_scores(repo_path)

    def process_chunk(self, file_path, output_folder, mapping):
        try:
            output = self.get_output(file_path).model_dump_json(indent=2)
            output_file_path = os.path.join(output_folder, f"{os.path.splitext(os.path.basename(file_path))[0]}.json")
            
            with open(output_file_path, "w", encoding="utf-8") as f:
                f.write(output)
            
            mapping[file_path] = output_file_path
        except Exception as e:
            self.logger.error(f"Error processing file {file_path}: {str(e)}")

    def final_scores(self, repo_path):
        if type(self.response_model) == CodeDescriptorModel:
            return
        directory = os.path.join(repo_path, "output_data")
        score_aggregation = defaultdict(int)
        files = 0

        for filename in os.listdir(directory):
            if filename.endswith('.json'):
                with open(os.path.join(directory, filename), 'r') as file:
                    data = json.load(file)

                for key, value in data.items():
                    if isinstance(value, dict) and 'score' in value:
                        score_aggregation[key] += value['score']
                files += 1

        for category in score_aggregation:
            score_aggregation[category] = round(score_aggregation[category] / files)

        output_data = {"scores_by_category": dict(score_aggregation)}
        output_file = os.path.join(repo_path, "output_data/scores_summary.json")

        with open(output_file, 'w') as file:
            json.dump(output_data, file, indent=2)

        self.logger.info(f"Scores summary saved to: {output_file}")

if __name__ == "__main__":
    repo_path = "./cloned_repos/regit"
    code_analyser = CodeAnalyser("code_security")
    code_analyser.process_repo(repo_path)