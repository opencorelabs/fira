from ninja import Router, Schema
from ninja_jwt.authentication import JWTAuth
from .models import User

router = Router(tags=['Users'])

class EmailPasswordRegistration(Schema):
    full_name: str
    email_address: str
    password: str


class UserResponse(Schema):
    email: str
    full_name: str
    verified: bool


class InputError(Schema):
    fields: dict[str, list[str]]  # map fields to list of errors
    message: str


@router.post('/register', response={201: UserResponse, 400: InputError})
def create_user(request, data: EmailPasswordRegistration):
    if User.objects.filter(email=data.email_address).exists():
        return 400, InputError(
            fields={'email_address': ['Email address already in use']},
            message='Email address already in use'
        )
    user = User.objects.create_user(
        email=data.email_address,
        password=data.password,
        full_name=data.full_name,
        verified=True
    )
    return 201, user

@router.get('/me', response=UserResponse, auth=JWTAuth())
def get_me(request):
    user = request.auth
    return user
