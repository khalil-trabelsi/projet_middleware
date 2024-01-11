import json
import requests
from sqlalchemy import exc
from marshmallow import EXCLUDE
from flask_login import current_user
import src.services.songs as songs_service
from src.schemas.rating import BaseRatingSchema, RatingSchema
from src.models.http_exceptions import *



ratings_url = "http://localhost:4200/ratings"  



def create_rating(newRating):
    songs, status_code = songs_service.get_song(newRating["song_id"])
  

    if status_code != 200:
        return songs, status_code
    print(newRating)
    response = requests.request(method="POST", url=ratings_url, json=newRating)
    if response.status_code != 201:
        return response.json(), response.status_code


    return response.json(), response.status_code

def get_rating(id):
    response = requests.request(method="GET", url=ratings_url+"/"+id)
    return response.json(), response.status_code

def update_rating(id, rating_update):  
    rating, status_code = get_rating(id)
    if status_code != 200:
        return rating, status_code
    
  

    response = requests.request(method="PUT", url=ratings_url+"/"+id, json=rating_update)
    if response.status_code != 200:
            return response.json(), response.status_code

    return response.json(), response.status_code

def delete_rating(id):
    rating, status_code = get_rating(id)
    if status_code != 200:
        return rating, status_code
    
   
    
    response = requests.request(method="DELETE", url=ratings_url+"/"+id)
    if response.status_code == 204:
        return "Rating supprimée avec succès", 204
    else:
        return response.json(), response.status_code

def get_ratings():
    response = requests.request(method="GET", url=ratings_url)
    return response.json(), response.status_code