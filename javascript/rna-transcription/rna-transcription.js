class DnaTranscriber {
  constructor(props) {
    this.DNA_MAPPING = new Map([
      ["C", "G"],
      ["G", "C"],
      ["A", "U"],
      ["T", "A"]
    ]);
  }
  toRna(seq) {
    return seq.split("").map(s => this.DNA_MAPPING.get(s)).join("");
  }
}

module.exports = DnaTranscriber;
