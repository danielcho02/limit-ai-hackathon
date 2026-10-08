import Hello from "./Components/Hello.js";

export default class App {
  constructor($target) {
    const hello = new Hello({
      $target,
    });
  }
}
