from django.contrib import admin
from .models import Account, AccountType, AccountTypeSubType

admin.site.register(Account)
admin.site.register(AccountType)
admin.site.register(AccountTypeSubType)