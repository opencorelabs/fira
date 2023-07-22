from typing import List
from django.shortcuts import get_object_or_404
from django.db.models import Q
from ninja import Schema, Router
from ninja_jwt.authentication import JWTAuth
from .models import Account, AccountType, AccountTypeSubType

router = Router(tags=['Accounts'])

class AccountRequest(Schema):
    name: str
    type_id: str
    subtype_id: str

class AccountResponse(Schema):
    name: str
    id: int
    type_id: str
    subtype_id: str

@router.post("/accounts", response=AccountResponse, auth=JWTAuth())
def create_account(request, payload: AccountRequest):
    at = get_object_or_404(AccountType, id=payload.type_id)
    ast = get_object_or_404(AccountTypeSubType, id=payload.subtype_id)
    account = Account.objects.create(user=request.user,name=payload.name,type=at,subtype=ast)
    return account

@router.get("/accounts/", response=List[AccountResponse], auth=JWTAuth())
def get_accounts(request):
    user = request.auth
    accounts = user.account_set.all()
    return accounts

@router.get("/accounts/{account_id}", response=AccountResponse, auth=JWTAuth())
def get_account(request, account_id: int):
    user = request.auth
    account = get_object_or_404(Account, Q(user_id=user.id) & Q(id=account_id))
    return account

@router.put("/accounts/{account_id}", response=AccountResponse, auth=JWTAuth())
def update_account(request, account_id: int, payload: AccountRequest):
    user = request.auth
    account = get_object_or_404(Account, Q(user_id=user.id) & Q(id=account_id))
    for attr, value in payload.dict().items():
        setattr(account, attr, value)
    account.save()
    return account

@router.delete("/accounts/", response={204: None}, auth=JWTAuth())
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