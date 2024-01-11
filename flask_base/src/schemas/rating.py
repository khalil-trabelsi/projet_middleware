from marshmallow import Schema, fields, validates_schema, ValidationError


#shéma rating de sortie (renvoyé au front)

class RatingSchema(Schema):
    
        id = fields.Integer(description="int")
        user_id = fields.String(description="User id")
        music_id = fields.String(description="Music id")
        content = fields.String(description="Content")
        date = fields.DateTime(description="Date")
        rating = fields.Float(description="Rating")
    
        @staticmethod
        def is_empty(obj):
            return (not obj.get("id") or obj.get("id") == "") and \
                (not obj.get("user_id") or obj.get("user_id") == "") and \
                (not obj.get("music_id") or obj.get("music_id") == "") and \
                (not obj.get("content") or obj.get("content") == "") and \
                (not obj.get("date") or obj.get("date") == "") and \
                (not obj.get("rating") or obj.get("rating") == "")
        
class BaseRatingSchema(Schema):
    user_id = fields.String(description="User id")
    music_id = fields.String(description="Music id")
    content = fields.String(description="Content")
    date = fields.DateTime(description="Date")
    rating = fields.Float(description="Rating")

# Schéma rating de modification (content, date, rating)
class RatingUpdateSchema(BaseRatingSchema):
    # permet de définir dans quelles conditions le schéma est validé ou nom
    @validates_schema
    def validates_schemas(self, data, **kwargs):
        if not (("content" in data and data["content"] != "") or
                ("date" in data and data["date"] != "") or
                ("rating" in data and data["rating"] != "")):
            raise ValidationError("at least one of ['content','date','rating'] must be specified")
        