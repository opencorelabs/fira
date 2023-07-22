from django.db import models
from core.models import TimeStampedModel
from accounts.models import Account

class Balance(TimeStampedModel):
    account = models.ForeignKey(Account, on_delete=models.CASCADE)
    balance = models.DecimalField(max_digits=255, decimal_places=31)
    effective_datetime = models.DateTimeField(null=False)
    
    class Meta:
        constraints = [
            models.UniqueConstraint(fields=['account', 'effective_datetime'], name='unique effective datetimes')
        ]

    def __str__(self):
        return self.account.user.email + ': ' + self.account.name + '-' + str(self.effective_datetime)