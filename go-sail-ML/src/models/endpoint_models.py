#models.py
from pydantic import BaseModel
from typing import Dict, List, Optional, Union, Any

class AnalysisRequest(BaseModel):
    analysis_type: str

class AnalysisResponse(BaseModel):
    analysis_type: str
    results: Dict[str, Any]

class Init(BaseModel):
    url: str