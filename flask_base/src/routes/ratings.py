import json
from flask import Blueprint, request
from flask_login import login_required
from marshmallow import ValidationError

from src.models.http_exceptions import *
from src.schemas.rating import RatingUpdateSchema
from src.schemas.errors import *
import src.services.ratings as ratings_service

  
ratings = Blueprint(name="ratings", import_name=__name__)

@ratings.route('/', methods=['POST'])
@login_required
def post_rating():
    """
    ---
    post:
      description: Creating a rating
      requestBody:
        required: true
        content:
            application/json:
                schema: RatingCreate
      responses:
        '201':
          description: Created
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - ratings
    """
    print( request.json)
    return ratings_service.create_rating(request.json)

@ratings.route('/<id>', methods=['GET'])
@login_required
def get_rating(id):

    """
    ---
    get:
      description: Getting a rating
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - ratings
    """
    return ratings_service.get_rating(id)

@ratings.route('/<id>', methods=['PUT'])
@login_required
def put_rating(id):
    """
    ---
    put:
      description: Updating a rating
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      requestBody:
        required: true
        content:
            application/json:
                schema: RatingUpdate
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Rating
            application/yaml:
              schema: Rating
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
        '422':
          description: Unprocessable entity
          content:
            application/json:
              schema: UnprocessableEntity
            application/yaml:
              schema: UnprocessableEntity
      tags:
          - ratings
    """
    return ratings_service.update_rating(id, request.json)
@ratings.route('/<id>', methods=['DELETE'])
@login_required
def delete_rating(id):
    """
    ---
    delete:
      description: Deleting a rating
      parameters:
        - in: path
          name: id
          schema:
            type: uuidv4
          required: true
          description: UUID of rating id
      responses:
        '204':
          description: No content
        '401':
          description: Unauthorized
          content:
            application/json:
              schema: Unauthorized
            application/yaml:
              schema: Unauthorized
        '404':
          description: Not found
          content:
            application/json:
              schema: NotFound
            application/yaml:
              schema: NotFound
      tags:
          - ratings
    """
    return ratings_service.delete_rating(id)


@ratings.route('/', methods=['GET'])
@login_required
def get_ratings():
    """
    ---
    get:
      description: Getting all ratings
      responses:
        '200':
          description: Ok
          content:
            application/json:
              schema: Ratings
            application/yaml:
              schema: Ratings
        '401':

            description: Unauthorized
            content:
                application/json:
                    schema: Unauthorized
                application/yaml:
                    schema: Unauthorized
    """
    return ratings_service.get_ratings()