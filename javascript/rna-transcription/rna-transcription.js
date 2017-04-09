const DNA_MAPPING = new Map([["C", "G"], ["G", "C"], ["A", "U"], ["T", "A"]]);
class DnaTranscriber {
  constructor(props) {}
  toRna(seq) {
    return seq.split("").map(s => DNA_MAPPING.get(s)).join("");
  }
}

module.exports = DnaTranscriber;
