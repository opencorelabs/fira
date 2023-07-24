from ninja import Router, Schema
from .models import User

router = Router(tags=['Users'])


class EmailPasswordRegistration(Schema):
    full_name: str
    email_address: str
    password: str
    verification_url: str


class EmailVerification(Schema):
    resend_to_email: str
    verification_code: str


class EmailPasswordLogin(Schema):
    email_address: str
    password: str


class Token(Schema):
    jwt: str


class Account(Schema):
    email_address: str
    full_name: str
    verified: bool


class InputError(Schema):
    fields: dict[str, list[str]]  # map fields to list of errors
    message: str


@router.post('/register', response={201: Account, 400: InputError})
def create_user(request, data: EmailPasswordRegistration):
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


@router.post('/verify', response={200: Account, 400: InputError})
def verify_user(request, data: EmailVerification):
    user = User.objects.filter(verification_code=data.verification_code).first()
    if user is None:
        return 400, InputError(
            fields={'verification_code': ['Invalid verification code']},
            message='Invalid verification code'
        )
    user.verified = True
    user.save()
    return 200, Account(
        email_address=user.email,
        full_name=user.full_name,
        verified=user.verified
    )


@router.post('/login', response={200: Token, 400: InputError})
def login(request, data: EmailPasswordLogin):
    user = User.objects.filter(email=data.email_address).first()
    if user is None:
        return 400, InputError(
            fields={'email_address': ['Email address not found']},
            message='Email address not found'
        )
    if not user.check_password(data.password):
        return 400, InputError(
            fields={'password': ['Incorrect password']},
            message='Incorrect password'
        )
    return 200, Token(jwt=user.get_jwt())


@router.get('/me', response={200: Account})
def get_me(request):
    user = request.auth
    return 200, Account(
        email_address=user.email,
        full_name=user.full_name,
        verified=user.verified
    )
