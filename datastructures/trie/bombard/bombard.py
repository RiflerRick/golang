import random
import string
import requests
import traceback
import sys
import threading


def random_char(y):
    return ''.join(random.choice(string.ascii_letters) for x in range(y))


class CompareAndSet(object):
    def __init__(self, key, init_value, runs, frequency):
        """
        key: key to test with
        init_value: initial value can be a number such as 1
        runs: number of runs to test it with
        frequency: after how many runs, a compare will be done
        """
        self._key = key
        if type(init_value) != int:
            raise Exception("init_value must be an integer")
        self._init_value = init_value
        self._runs = runs
        self._frequency = frequency
        self._successful_compares = 0

    @property
    def successful_compares(self):
        return self._successful_compares

    def _update(self, value):
        response = requests.get(
            "http://127.0.0.1:1337/put/?key={}&value={}".format(self._key, value))
        if response.status_code != 200:
            raise Exception
        print("updating with {}".format(value))

    def _compare(self, expected_value):
        response = requests.get(
            "http://127.0.0.1:1337/get/?key={}".format(self._key))
        if response.status_code != 200:
            raise Exception
        # import ipdb
        # ipdb.set_trace()
        if str(response.text) == str(expected_value):
            return True, None
        return False, response.text

    def run(self):
        try:
            self._update(self._init_value)
            print("first update successful")
            expected_value = self._init_value
            for i in range(self._runs):

                if i % self._frequency == 0:
                    print("comaparing")
                    found, response_text = self._compare(expected_value)
                    if found:
                        print("compare successful")
                        self._successful_compares += 1
                    else:
                        print("compare failed: expected: {} but found: {}".format(
                            expected_value, response_text))
                    print("reinitializing value")
                    self._update(self._init_value)
                    expected_value = self._init_value
                    continue
                print("updating value")
                expected_value += 1
                self._update(expected_value)
        except Exception:
            print(traceback.format_exc())


def put(samples, key_size, value_size, randomize):
    key_list = []
    for i in range(samples):
        key = random_char(key_size)
        value = random_char(value_size)
        key_list.append(key)
        if not randomize:
            response = requests.get(
                "http://127.0.0.1:1337/put/?key={}&value={}".format("ilapahsi", "alice"))
            if response.status_code == 200:
                print(response.text)
            else:
                print("ERROR: error from server: {}".format(str(response)))
            continue
        response = requests.get(
            "http://127.0.0.1:1337/put/?key={}&value={}".format(key, value))
        if response.status_code == 200:
            print(response.text)
        else:
            print("ERROR: error from server: {}".format(str(response)))
    return key_list


def get(key_list, randomize):
    for i in range(key_list):
        key = key_list[i]
        if not randomize:
            response = requests.get(
                "http://127.0.0.1:1337/get/?key={}".format("ilapahsi"))
            if response.status_code == 200:
                print(response.text)
            else:
                print("ERROR: error from server: {}".format(str(response)))
            continue
        response = requests.get(
            "http://127.0.0.1:1337/get/?key={}".format(key))
        if response.status_code == 200:
            print(response.text)
        else:
            print("ERROR: error from server: {}".format(str(response)))


if __name__ == "__main__":
    if len(sys.argv) != 2:
        raise Exception(
            "expecting 1 command line argument: either --compare-and-set or --bombard")
    arg = sys.argv[1]
    if arg == "--compare-and-set":
        c = CompareAndSet("ilapahsi", 1, 1000, 4)
        c.run()
        print("successful compares: {}".format(c.successful_compares))
    elif arg == "--bombard":
        samples = 1000
        key_size = 10
        value_size = 20
        randomize = True
        key_list = put(samples, key_size, value_size, randomize)
        get(key_list, randomize)
    else:
        raise Exception
