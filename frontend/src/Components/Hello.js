export default class Hello {
  constructor({ $target }) {
    this.section = document.createElement("section");
    $target.appendChild(this.section);
    this.render();
  }

  render() {
    this.section.innerHTML = "";
    const text = document.createElement("h1");
    text.innerText = "hello world!";
    this.section.appendChild(text);
  }
}
