from decimal import Decimal
from typing import List,Optional
from datetime import datetime
from django.shortcuts import get_object_or_404
from django.db.models import Q
from ninja import ModelSchema,Schema,FilterSchema,Query,Field
from ninja_extra.ordering import ordering, Ordering
from ninja_extra import api_controller, route, ControllerBase
from ninja_extra.pagination import paginate, PageNumberPaginationExtra, PaginatedResponseSchema
from ninja_jwt.authentication import JWTAuth
from .models import Balance
from accounts.models import Account

class BalanceFilter(FilterSchema):
    balance: Optional[Decimal]
    start_datetime: Optional[datetime] = Field(q='effective_datetime__gte') 
    end_datetime: Optional[datetime] = Field(q='effective_datetime__lte') 

class BalanceRequest(Schema):
    balance: Decimal
    effective_datetime: datetime

class BalanceResponse(ModelSchema):
    class Config:
        model = Balance
        model_fields = ['balance', 'id', 'effective_datetime']


@api_controller('/accounts/{account_id}',tags=['Balances'], auth=JWTAuth())
class BalanceController(ControllerBase):
    @route.post("/balances/", response=BalanceResponse )
    def create_balance(self, account_id: int, payload: BalanceRequest):
        user = self.context.request.auth
        account = get_object_or_404(Account, Q(user_id=user.id) & Q(id=account_id))
        balance = Balance.objects.create(account=account, balance= payload.balance, effective_datetime= payload.effective_datetime)
        return balance

    @route.get("/balances/", response=PaginatedResponseSchema[BalanceResponse])
    @paginate(PageNumberPaginationExtra, page_size=50)
    @ordering(Ordering, ordering_fields=['effective_datetime', 'balance'])
    def get_balances(self, account_id: int, filters: BalanceFilter = Query(...)):
        user = self.context.request.auth
        balances = Balance.objects.filter(Q(account__user_id=user.id) & Q(account_id=account_id))
        balances = filters.filter(balances)
        return balances

    @route.get("/balances/{balance_id}", response=BalanceResponse)
    def get_balance(self, account_id: int, balance_id: int):
        user = self.context.request.auth
        balance = get_object_or_404(Balance, Q(account__user_id=user.id) & Q(account_id=account_id) & Q(id=balance_id))
        return balance

    @route.put("/balances/{balance_id}", response=BalanceResponse)
    def update_balance(self, account_id: int, balance_id: int, payload: BalanceRequest):
        user = self.context.request.auth
        balance = get_object_or_404(Balance, Q(account__user_id=user.id) & Q(account_id=account_id) & Q(id=balance_id))
        for attr, value in payload.dict().items():
            setattr(balance, attr, value)
        balance.save()
        return balance

    @route.delete("/balances/", response={204: None})
    def delete_balances(self, account_id: int):
        user = self.context.request.auth
        balances = Balance.objects.filter(Q(account__user_id=user.id) & Q(account_id=account_id))
        balances.delete()
        return 204

    @route.delete("/balances/{balance_id}", response={204: None})
    def delete_balance(self, account_id: int, balance_id: int):
        user = self.context.request.auth
        balance = get_object_or_404(Balance, Q(account__user_id=user.id) & Q(account_id=account_id) & Q(id=balance_id))
        balance.delete()
        return 204