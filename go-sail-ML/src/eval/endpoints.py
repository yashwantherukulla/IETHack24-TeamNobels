from fastapi import FastAPI, HTTPException
from typing import Dict
from src.models.endpoint_models import AnalysisRequest, AnalysisResponse, Init
from src.eval.git_handler import GitHandler
from src.eval.chunker import ChunkExtractor
from src.eval.code_analyser import CodeAnalyser
import os
import json


app = FastAPI()
BASE_PATH = "./cloned_repos"
REPO_PATH = ''


def read_output_data(repo_path):
    output_data_path = os.path.join(repo_path, "output_data")
    result = {}
    
    for root, _, files in os.walk(output_data_path):
        for file in files:
            if file.endswith(".json"):
                file_path = os.path.join(root, file)
                with open(file_path, 'r') as f:
                    chunk_data = json.load(f)
                
                # Extract the original file path from the JSON file name
                relative_path = file.replace('#', '/').rsplit('-', 1)[0]
                if relative_path not in result:
                    result[relative_path] = []
                result[relative_path].append(chunk_data)
    
    return result


@app.post("/init")
def init(data: Init) -> Dict[str, str]:
    git_handler = GitHandler()
    git_handler = GitHandler(BASE_PATH)
    git_handler.clone_repository(data.url)
    repo_name = os.path.splitext(data.url.split('/')[-1])[0]
    global REPO_PATH
    REPO_PATH = os.path.join(BASE_PATH, repo_name)
    chunk_extractor = ChunkExtractor()
    chunk_extractor.processRepos(BASE_PATH)
    return {"message": "Repository initialized."}

@app.post("/analyze/")
def analyze_code(data: AnalysisRequest):
    code_analyser = CodeAnalyser(data.analysis_type)
    code_analyser.process_repo(REPO_PATH)
    output_data = read_output_data(REPO_PATH)
    
    return {"message": "Repository analyzed.", "output_data": output_data}


if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)

