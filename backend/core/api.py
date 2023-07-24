from typing import List
from django.shortcuts import get_object_or_404
from django.db.models import Q
from ninja import Schema, Router
from .models import Account, AccountType, AccountTypeSubType

router = Router(tags=['Core'])

class AccountRequest(Schema):
    name: str
    type_id: str
    subtype_id: str

class AccountResponse(Schema):
    name: str
    id: int
    type_id: str
    subtype_id: str

@router.post("/core", response=AccountResponse)
def create_account(request, payload: AccountRequest):
    at = get_object_or_404(AccountType, id=payload.type_id)
    ast = get_object_or_404(AccountTypeSubType, id=payload.subtype_id)
    account = Account.objects.create(user=request.user,name=payload.name,type=at,subtype=ast)
    return account

@router.get("/core/", response=List[AccountResponse])
def get_accounts(request):
    user = request.auth
    accounts = user.account_set.all()
    return accounts

@router.get("/core/{account_id}", response=AccountResponse)
def get_account(request, account_id: int):
    user = request.auth
    account = get_object_or_404(Account, Q(user_id=user.id) & Q(id=account_id))
    return account

@router.put("/core/{account_id}", response=AccountResponse)
def update_account(request, account_id: int, payload: AccountRequest):
    user = request.auth
    account = get_object_or_404(Account, Q(user_id=user.id) & Q(id=account_id))
    for attr, value in payload.dict().items():
        setattr(account, attr, value)
    account.save()
    return account

@router.delete("/core/", response={204: None})
def delete_accounts(request):
    user = request.auth
    user.account_set.all().delete()
    return 204


@router.delete("/core/{account_id}", response={204: None})
def delete_account(request, account_id: int):
    user = request.auth
    account = get_object_or_404(Account, Q(user_id=user.id) & Q(id=account_id))
    account.delete()
    return 204