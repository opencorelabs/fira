from typing import List
from django.shortcuts import get_object_or_404
from django.db.models import Q
from ninja import Field,Schema, Router
from ninja_jwt.authentication import JWTAuth
from .models import Account, AccountType

router = Router(tags=['Manual Accounts'])

class AccountInputSchema(Schema):
    name: str
    type_id: int

class AccountOutputSchema(Schema):
    id: int
    name: str
    type_id: int

class AccountTypeOutputSchema(Schema):
    id: int 
    name: str
    classification: str

@router.get("/accounttypes", response=List[AccountTypeOutputSchema])
def get_account_types(request):
    ats = AccountType.objects.all()
    return ats

@router.post("/accounts", response=AccountOutputSchema, auth=JWTAuth())
def create_account(request, payload: AccountInputSchema):
    t = get_object_or_404(AccountType, id=payload.type_id)
    account = Account.objects.create(user=request.user,name=payload.name,type=t)
    return account

@router.get("/accounts", response=List[AccountOutputSchema], auth=JWTAuth())
def get_accounts(request):
    user = request.auth
    accounts = user.account_set.all()
    return accounts

@router.get("/accounts/{account_id}", response=AccountOutputSchema, auth=JWTAuth())
def get_account(request, account_id: int):
    user = request.auth
    account = get_object_or_404(Account, Q(user_id=user.id) & Q(id=account_id))
    return account

@router.put("/accounts/{account_id}", response=AccountOutputSchema, auth=JWTAuth())
def update_account(request, account_id: int, payload: AccountInputSchema):
    user = request.auth
    account = get_object_or_404(Account, Q(user_id=user.id) & Q(id=account_id))
    for attr, value in payload.dict().items():
        setattr(account, attr, value)
    account.save()
    return account

@router.delete("/accounts", response={204: None}, auth=JWTAuth())
def delete_accounts(request):
    user = request.auth
    user.account_set.all().delete()
    return 204


@router.delete("/accounts/{account_id}", response={204: None}, auth=JWTAuth())
def delete_account(request, account_id: int):
    user = request.auth
    account = get_object_or_404(Account, Q(user_id=user.id) & Q(id=account_id))
    account.delete()
    return 204
