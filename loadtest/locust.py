from locust import HttpUser, TaskSet, task
from bs4 import BeautifulSoup
import time
import random

class WebsiteUser(HttpUser):

    def on_start(self):
        self.cache = set()

    def hit_and_sleep(self, resource, min_sleep, max_sleep):
        response_text = ''
        with self.client.get(resource, catch_response=True) as response:
            response_text = response.text
        # Download static resources if not in this user's cache already.
        """
        if not resource in self.cache:
            soup = BeautifulSoup(response_text, features="html.parser")
            # Download static images.
            for elem in soup.find_all('img'):
                self.client.get(elem.get('src'))
            # Download CSS.
            for elem in soup.find_all('link'):
                self.client.get(elem.get('href'))
            # Download JavaScript.
            for elem in soup.find_all('script'):
                self.client.get(elem.get('src'))
            # Mark the resource as cached for this user.
            self.cache.add(resource)
        """
        # Sleep a random amount in the given range.
        self.sleep(min_sleep, max_sleep)

    def sleep(self, min, max):
        r = random.uniform(min, max)
        time.sleep(r)

    @task(6)
    def home(self):
        self.hit_and_sleep('/home', 3.0, 15.0)

    @task(4)
    def jobs(self):
        self.hit_and_sleep('/jobs', 3.0, 20.0)

    @task(2)
    class About(TaskSet):

        @task(6)
        def about(self):
            self.user.hit_and_sleep('/about', 3.0, 15.0)


        @task(3)
        def offices(self):
            self.user.hit_and_sleep('/about/offices', 3.0, 15.0)


        @task(1)
        def legals(self):
            self.user.hit_and_sleep('/about/legals', 4.0, 40.0)

        @task(5)
        def back(self):
            self.interrupt()