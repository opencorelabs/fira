from ninja import Router, Schema
from .models import User

router = Router(tags=['Accounts'])


class EmailPasswordRegistration(Schema):
    full_name: str
    email_address: str
    password: str


class Account(Schema):
    email_address: str
    full_name: str
    verified: bool


class InputError(Schema):
    fields: dict[str, list[str]]  # map fields to list of errors
    message: str


@router.post('/register', response={201: Account, 400: InputError})
def create_account(request, data: EmailPasswordRegistration):
    if User.objects.filter(email=data.email_address).exists():
        return 400, InputError(
            fields={'email_address': ['Email address already in use']},
            message='Email address already in use'
        )
    user = User.objects.create_user(
        email=data.email_address,
        password=data.password,
        full_name=data.full_name
    )
    return 201, Account(
        email_address=user.email,
        full_name=user.full_name,
        verified=user.verified
    )
