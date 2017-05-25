class PhoneNumber {
  constructor(number) {
    this.num = this.validate(number);
  }

  validate(number) {
    const INVALID_NUMBER = "0000000000";
    if (!number) {
      return INVALID_NUMBER;
    }
    number = number.replace(/[\.\(\)\- ]/g, "");
    if (this.startsWith11(number)) {
      // if its starts with 11 then allow
      return number.slice(1, number.length);
    } else if (number.length != 10) {
      return INVALID_NUMBER;
    } else {
      return number;
    }
  }
  number() {
    return this.num;
  }
  areaCode() {
    return this.num.slice(0, 3);
  }

  startsWith11(n) {
    return n.length == 11 && n.slice(0, 2) == "11";
  }
  toString() {
    const firstThree = this.num.slice(3, 6);
    const lastFour = this.num.slice(6, 10);
    return `(${this.areaCode()}) ${firstThree}-${lastFour}`;
  }
}

module.exports = PhoneNumber;
