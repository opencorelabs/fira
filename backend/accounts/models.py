from django.db import models
from core.models import TimeStampedModel
from users.models import User

class AccountType(TimeStampedModel):
    id = models.CharField(max_length=15, primary_key=True)
    name = models.CharField(max_length=31,null=False, blank=False)
    is_asset = models.BooleanField(null=False)

    class Meta:
        constraints = [
            models.UniqueConstraint(fields=['name'], name='unique account type name')
        ]

    def __str__(self):
        return self.name

class AccountTypeSubType(TimeStampedModel):
    id = models.CharField(max_length=15, primary_key=True)
    name = models.CharField(max_length=31, null=False, blank=False)
    account_type = models.ManyToManyField(AccountType)

    class Meta:
        constraints = [
            models.UniqueConstraint(fields=['name'], name='unique account sub type name')
        ]

    def __str__(self):
        return self.name

class Account(TimeStampedModel):
    user = models.ForeignKey(User, on_delete=models.CASCADE, null=False)
    name = models.CharField(max_length=255, null=False, blank=False) 
    type = models.ForeignKey(AccountType, on_delete=models.RESTRICT, null=False)
    subtype = models.ForeignKey(AccountTypeSubType, on_delete=models.RESTRICT, null=False)

    class Meta:
        constraints = [
            models.UniqueConstraint(fields=['user', 'name'], name='unique user account names')
        ]
      
    def __str__(self):
        return self.user.email + ': ' + self.name