from django.contrib import admin
from .models import Account, AccountType

admin.site.register(Account)
admin.site.register(AccountType)