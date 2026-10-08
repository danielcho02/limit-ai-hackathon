import { makeObservable, observable, action, computed } from "mobx";

class TabStore {
  screenMode = "Home";
  launched = true;

  constructor() {
    makeObservable(this, {
      screenMode: observable,
      changeScreenMode: action,
    });

    makeObservable(this, {
      launched: observable,
      setFirstLaunch: action,
    });
  }

  changeScreenMode(a) {
    this.screenMode = a;
  }

  setFirstLaunch(a) {
    this.launched = a;
  }
}

export const tabStore = new TabStore();
