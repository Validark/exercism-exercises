class Pangram {
  constructor(phrase) {
    this.ALPHA_CONST = 97;
    if (phrase === null) {
      this.phrase = "";
    }
    this.phrase = phrase.toLowerCase();
  }

  isPangram() {
    const check = new Array(25).fill(0);
    this.phrase.split("").map(letter => {
      const code = letter.charCodeAt() - this.ALPHA_CONST;
      if (code >= 0 && code < 26) {
        check[code] = 1;
      }
    });
    return check.reduce((a, b) => a + b) == 26;
  }
}

module.exports = Pangram;
