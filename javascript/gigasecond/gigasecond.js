class Gigasecond {
  constructor(time) {
    this.GIG = 1e12;
    this.time = time;
  }
  date() {
    return new Date(this.time.getTime() + this.GIG);
  }
}

module.exports = Gigasecond;
