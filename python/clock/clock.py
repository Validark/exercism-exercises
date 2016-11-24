import datetime

class Clock():
    def __init__(self, hour=0, minute=0):
        self.t = datetime.timedelta(hours=hour,minutes=minute) + datetime.datetime.utcfromtimestamp(0)

    def __repr__(self):
        return self.t.time().strftime('%H:%M')

    def __eq__(self, other):
        """Override the default Equals behavior"""
        if isinstance(other, self.__class__):
            return self.__repr__() == other.__repr__()
        return NotImplemented

    def add(self, minute):
        self.t = self.t + datetime.timedelta(minutes=minute)
        return self.__repr__()