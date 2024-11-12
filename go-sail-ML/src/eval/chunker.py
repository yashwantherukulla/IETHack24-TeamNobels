#chunker.py
import os
import json
import logging
from pathlib import Path
from llama_index.core.node_parser import CodeSplitter
from llama_index.readers.file import FlatReader
from ..eval.languages import languageExtensions

class ChunkExtractor:
    def __init__(self):
        self.logger = logging.getLogger(__name__)
        self.logger.setLevel(logging.INFO)
        handler = logging.StreamHandler()
        handler.setFormatter(logging.Formatter('%(asctime)s - %(name)s - %(levelname)s - %(message)s'))
        self.logger.addHandler(handler)

    def detectLanguage(self, filePath):
        extension = Path(filePath).suffix[1:].lower()
        print(languageExtensions.get(extension, 'unknown'), "\n\n\n")
        return languageExtensions.get(extension, 'unknown')

    def processRepos(self, root_folder):
        root_path = Path(root_folder)
        for repo_path in root_path.iterdir():
            if repo_path.is_dir():
                self.processRepo(repo_path)

    def processRepo(self, repo_path):
        chunk_folder = repo_path / "chunk_data"
        chunk_folder.mkdir(exist_ok=True)
        
        for file_path in repo_path.rglob('*'):
            if file_path.is_file() and (str(chunk_folder) not in str(file_path)):
                self.processFile(file_path, chunk_folder)

    def processFile(self, file_path, chunk_folder):
        try:
            language = self.detectLanguage(file_path)
            if language == 'unknown':
                self.logger.info(f"Skipping file with unknown language: {file_path}")
                return
            
            document = FlatReader().load_data(file_path)
            
            code_splitter = CodeSplitter(chunk_lines=1000, language=language, max_chars=6000)
            nodes = code_splitter.get_nodes_from_documents(document)
            
            # Create relative path string for naming
            rel_path = str(file_path)
            # Replace path separators with dashes
            file_identifier = rel_path.replace('/', '#').replace('\\', '##')
            
            for i, node in enumerate(nodes, 1):
                chunk_file_path = chunk_folder / f"{file_identifier}-{i}.json"
                chunk_file_path.write_text(node.to_json(), encoding="utf-8")
            
            self.logger.info(f"Processed file: {file_path}")
                
        except Exception as e:
            self.logger.error(f"Error processing file {file_path}: {str(e)}", exc_info=True)

if __name__ == "__main__":
    base_path = Path("./cloned_repos")
    chunk_extractor = ChunkExtractor()
    chunk_extractor.processRepos(base_path)

    