import { tabStore } from "./store";
import { useEffect, useRef, useState } from "react";
import Home from "./Home";
import Market from "./Market";
import { observer } from "mobx-react-lite";
import OnBoarding from "./onboarding";

function Navigation() {
  console.log("Navigation render", tabStore.screenMode);
  return (
    <div
      style={{
        paddingLeft: 40,
        paddingRight: 40,
        paddingTop: 20,
        paddingBottom: 20,
        left: 0,
        bottom: 0,
        position: "fixed",
        background: "white",
        boxShadow: "0px -4px 12px rgba(0, 0, 0, 0.25)",
        borderTopLeftRadius: 24,
        borderTopRightRadius: 24,
        flexDirection: "column",
        justifyContent: "flex-start",
        alignItems: "flex-start",
        gap: 10,
        display: "inline-flex",
      }}
    >
      <div
        style={{
          alignSelf: "start",
          justifyContent: "flex-start",
          alignItems: "center",
          gap: 57,
          display: "inline-flex",
        }}
      >
        <div
          style={{ width: 36, height: 52, position: "relative" }}
          onClick={() => {
            tabStore.changeScreenMode("Home");
          }}
        >
          <img
            src={
              tabStore.screenMode === "Home"
                ? require("./assets/homeactiveicn.png")
                : require("./assets/homeicn.png")
            }
            style={{ width: 36, height: 36 }}
          />
          {/* <div
            style={{
              width: 36,
              height: 36,
              left: 0,
              top: 0,
              position: "absolute",
              overflow: "hidden",
            }}
          >
            <div
              style={{
                width: 24,
                height: 27,
                left: 6,
                top: 4.5,
                position: "absolute",
                background: "#DD0200",
              }}
            />
          </div> */}
          <div
            style={{
              width: 36,
              left: 0,
              top: 36,
              position: "absolute",
              textAlign: "center",
              color: tabStore.screenMode === "Home" ? "#DD0200" : "black",
              fontSize: 11,
              fontFamily: "Noto Sans KR",
              fontWeight: "900",
              wordWrap: "break-word",
            }}
          >
            Home
          </div>
        </div>
        <div
          style={{ width: 42, height: 51.5, position: "relative" }}
          onClick={() => {
            tabStore.changeScreenMode("Market");
          }}
        >
          <img
            src={
              tabStore.screenMode === "Market"
                ? require("./assets/marketactiveicn.png")
                : require("./assets/marketicn.png")
            }
            style={{ width: 36, height: 36 }}
          />
          <div
            style={{
              width: 36,
              left: 0,
              top: 35.5,
              position: "absolute",
              textAlign: "center",
              color: tabStore.screenMode === "Market" ? "#DD0200" : "black",
              fontSize: 11,
              fontFamily: "Noto Sans KR",
              fontWeight: "900",
              wordWrap: "break-word",
            }}
          >
            Market
          </div>
        </div>
        <div
          style={{ width: 36, height: 52, position: "relative" }}
          onClick={() => {
            tabStore.changeScreenMode("Chat");
          }}
        >
          <img
            src={
              tabStore.screenMode === "Chat"
                ? require("./assets/chatactiveicn.png")
                : require("./assets/chaticn.png")
            }
            style={{ width: 36, height: 36 }}
          />

          <div
            style={{
              width: 36,
              left: 0,
              top: 36,
              position: "absolute",
              textAlign: "center",
              color: tabStore.screenMode === "Chat" ? "#DD0200" : "black",
              fontSize: 11,
              fontFamily: "Noto Sans KR",
              fontWeight: "900",
              wordWrap: "break-word",
            }}
          >
            Chat
          </div>
        </div>
        <div
          style={{ width: 36, height: 52, position: "relative" }}
          onClick={() => {
            tabStore.changeScreenMode("Mypage");
          }}
        >
          <img
            src={
              tabStore.screenMode === "Mypage"
                ? require("./assets/mypageactiveicn.png")
                : require("./assets/mypageicn.png")
            }
            style={{ width: 36, height: 36 }}
          />
          <div
            style={{
              width: 36,
              left: 0,
              top: 36,
              position: "absolute",
              textAlign: "center",
              color: tabStore.screenMode === "Mypage" ? "#DD0200" : "black",
              fontSize: 11,
              fontFamily: "Noto Sans KR",
              fontWeight: "900",
              wordWrap: "break-word",
            }}
          >
            My
          </div>
        </div>
      </div>
    </div>
  );
}
const App = observer(() => {
  useEffect(() => {
    if (!tabStore.launched) return;

    const timer = setTimeout(() => {
      tabStore.setFirstLaunch(false);
    }, 2000);

    return () => clearTimeout(timer);
  }, [tabStore.launched]);

  if (tabStore.launched) {
    return <OnBoarding />;
  }

  return (
    <>
      {/* {tabStore.screenMode === "Home" && <Home />} */}
      {tabStore.screenMode === "Home" && (
        <img
          src={require("./assets/home.png")}
          style={{ width: "100%", height: "100%" }}
        />
      )}
      {tabStore.screenMode === "Market" && (
        <img
          src={require("./assets/market.png")}
          style={{ width: "100%", height: "100%" }}
        />
      )}
      {tabStore.screenMode === "Chat" && (
        <img
          src={require("./assets/chat.png")}
          style={{ width: "100%", height: "100%" }}
        />
      )}
      {tabStore.screenMode === "Mypage" && (
        <img
          src={require("./assets/mypage.png")}
          style={{ width: "100%", height: "100%" }}
        />
      )}
      <Navigation />
    </>
  );
});

export default App;
