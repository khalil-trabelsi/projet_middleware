from marshmallow import Schema, fields, validates_schema, ValidationError



class RatingSchema(Schema):
    
        id = fields.String(description="ID")
        user_id = fields.String(description="User id")
        song_id = fields.String(description="Song id")
        content = fields.String(description="Content")
        date = fields.DateTime(description="Date")
        rating = fields.Integer(description="Rating")
    
        @staticmethod
        def is_empty(obj):
            return (not obj.get("id") or obj.get("id") == "") and \
                (not obj.get("user_id") or obj.get("user_id") == "") and \
                (not obj.get("song_id") or obj.get("song_id") == "") and \
                (not obj.get("content") or obj.get("content") == "") and \
                (not obj.get("date") or obj.get("date") == "") and \
                (not obj.get("rating") or obj.get("rating") == "")
        
class BaseRatingSchema(Schema):
    user_id = fields.String(description="User id")
    song_id = fields.String(description="Song id")
    content = fields.String(description="Content")
    rating = fields.Integer(description="Rating")

class RatingUpdateSchema(BaseRatingSchema):
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("content" in data and data["content"] != "") or
                ("rating" in data and data["rating"] != "")):
            raise ValidationError("at least one of ['content','date','rating'] must be specified")
        