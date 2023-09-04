from django.db import models
from accounts.models import User
from core.models import TimeStampedModel

class AccountType(TimeStampedModel):
    name = models.CharField(max_length=31,null=False, blank=False)

    CLASSIFICATION_CHOICES = [
        ("Asset", "Asset"),
        ("Liability", "Liability"),
    ]
    classification = models.CharField(max_length=15,
                            choices=CLASSIFICATION_CHOICES,
                            null=False,
                            blank=False)
    
    class Meta:
        constraints = [
            models.UniqueConstraint(fields=['name', 'classification'], name='unique_account_type_name_classification')
        ]

    def __str__(self):
        return self.name

class Account(TimeStampedModel):
    user = models.ForeignKey(User, on_delete=models.CASCADE, null=False)
    name = models.CharField(max_length=255, null=False, blank=False) 
    type = models.ForeignKey(AccountType, on_delete=models.RESTRICT, null=False)

    class Meta:
        constraints = [
            models.UniqueConstraint(fields=['user', 'name'], name='unique_user_account_name')
        ]
      
    def __str__(self):
        return self.user.email + ': ' + self.name
    
    