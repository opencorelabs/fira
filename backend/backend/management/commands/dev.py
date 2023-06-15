import os
from django.core.management.base import BaseCommand
from django.core.management import call_command
from django.contrib.auth import get_user_model


class Command(BaseCommand):
    help = "Runserver and migrate"

    def handle(self, *args, **options):
        call_command('migrate')
        User = get_user_model()
        username = os.environ.setdefault('DJANGO_SUPERUSER_USERNAME', 'admin')
        password = os.environ.setdefault('DJANGO_SUPERUSER_PASSWORD', 'admin')
        email = os.environ.setdefault('DJANGO_SUPERUSER_EMAIL', 'test@test.net')
        if not User.objects.filter(username=username).exists():
            User.objects.create_superuser(username=username, password=password, email=email)
        call_command('runserver')
