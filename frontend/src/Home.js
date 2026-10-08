import logo from "./logo.svg";
import "./App.css";

import { observer } from "mobx-react-lite";
import { tabStore } from "./store";

const Home = observer(() => {
  return (
    <div className="App">
      <header className="App-header">
        <div
          style={{
            width: 393,
            height: 874,
            position: "relative",
            background: "white",
            overflow: "hidden",
            justifyContent: "flex-start",
            alignItems: "flex-start",
            gap: 10,
            display: "inline-flex",
          }}
        >
          <div
            style={{
              width: 393,
              height: 874,
              flexDirection: "column",
              justifyContent: "flex-start",
              alignItems: "flex-start",
              gap: 48,
              display: "inline-flex",
            }}
          >
            <div
              style={{
                width: 393,
                height: 338,
                paddingTop: 120,
                paddingBottom: 92,
                paddingLeft: 36,
                paddingRight: 37,
                background: "white",
                boxShadow: "0px 4px 50px rgba(32, 32, 32, 0.20)",
                borderBottomRightRadius: 36,
                borderBottomLeftRadius: 36,
                flexDirection: "column",
                justifyContent: "center",
                alignItems: "center",
                gap: 36,
                display: "flex",
              }}
            >
              <div
                style={{
                  width: 393,
                  paddingTop: 9,
                  paddingBottom: 9,
                  paddingLeft: 24,
                  paddingRight: 36,
                  overflow: "hidden",
                  justifyContent: "space-between",
                  alignItems: "center",
                  display: "inline-flex",
                }}
              >
                <div style={{ width: 338, height: 24, position: "relative" }}>
                  <div
                    style={{
                      width: 24,
                      height: 24,
                      left: 314.5,
                      top: 0,
                      position: "absolute",
                      overflow: "hidden",
                    }}
                  >
                    <div
                      style={{
                        width: 16,
                        height: 20,
                        left: 4,
                        top: 2,
                        position: "absolute",
                        background: "#202020",
                      }}
                    />
                  </div>
                </div>
              </div>
              <div
                style={{
                  width: 329,
                  flexDirection: "column",
                  justifyContent: "center",
                  alignItems: "center",
                  gap: 12,
                  display: "flex",
                }}
              >
                <img src={require("./assets/Frame 23.png")} />

                {/* <div
                  style={{
                    width: 329,
                    height: 95,
                    position: "relative",
                    overflow: "hidden",
                  }}
                >
                  <div
                    style={{
                      width: 42.36,
                      height: 46.12,
                      left: 0.35,
                      top: 2.47,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 22.02,
                      height: 45.57,
                      left: 44.19,
                      top: 10.71,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 37.09,
                      height: 71.72,
                      left: 66.95,
                      top: 10.44,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 25.54,
                      height: 66.51,
                      left: 103.24,
                      top: 15.48,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 41.48,
                      height: 13.22,
                      left: 120.41,
                      top: 64.13,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 18.13,
                      height: 17.45,
                      left: 168.66,
                      top: 66.51,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 25.81,
                      height: 47.14,
                      left: 167.15,
                      top: 16.01,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 72.52,
                      height: 45.81,
                      left: 195,
                      top: 5.19,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 81.57,
                      height: 52.56,
                      left: 191.2,
                      top: 29.97,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 28.71,
                      height: 24.28,
                      left: 268.31,
                      top: 0,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 45.74,
                      height: 21.04,
                      left: 270.08,
                      top: 13.78,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 57.18,
                      height: 23.97,
                      left: 266.06,
                      top: 22.09,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 38.99,
                      height: 48.66,
                      left: 126.7,
                      top: 15.18,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 46.64,
                      height: 30.62,
                      left: 270.69,
                      top: 42.08,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 329,
                      height: 84.39,
                      left: 0,
                      top: 10.61,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 5.19,
                      height: 4.18,
                      left: 301.38,
                      top: 50.16,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                  <div
                    style={{
                      width: 32.47,
                      height: 10.24,
                      left: 131.04,
                      top: 72.71,
                      position: "absolute",
                      background: "#DD0200",
                    }}
                  />
                </div> */}
                <div
                  style={{
                    alignSelf: "stretch",
                    height: 35,
                    textAlign: "center",
                    color: "var(--블랙, #202020)",
                    fontSize: 20,
                    fontFamily: "Noto Sans KR",
                    fontWeight: "700",
                    lineHeight: 36,
                    wordWrap: "break-word",
                  }}
                >
                  나도 몰랐던 내 재능의 쓸모
                </div>
              </div>
            </div>
            <div
              style={{
                alignSelf: "stretch",
                flexDirection: "column",
                justifyContent: "flex-start",
                alignItems: "flex-start",
                gap: 36,
                display: "flex",
              }}
            >
              <div
                style={{
                  alignSelf: "stretch",
                  height: 278,
                  position: "relative",
                }}
              >
                <div
                  style={{
                    width: 440,
                    left: 24,
                    top: 0,
                    position: "absolute",
                    flexDirection: "column",
                    justifyContent: "flex-start",
                    alignItems: "flex-start",
                    gap: 16,
                    display: "inline-flex",
                  }}
                >
                  <div
                    style={{
                      alignSelf: "stretch",
                      color: "black",
                      fontSize: 24,
                      fontFamily: "Noto Sans KR",
                      fontWeight: "900",
                      wordWrap: "break-word",
                    }}
                  >
                    @@@님을 위한 추천
                  </div>
                  <div
                    style={{
                      alignSelf: "stretch",
                      justifyContent: "flex-start",
                      alignItems: "center",
                      gap: 12,
                      display: "inline-flex",
                    }}
                  >
                    <div
                      data-속성-1="기본"
                      style={{
                        width: 136,
                        height: 207,
                        paddingTop: 12,
                        paddingBottom: 19,
                        paddingLeft: 8,
                        paddingRight: 8,
                        background: "white",
                        boxShadow: "0px 4px 20px rgba(32, 32, 32, 0.20)",
                        borderRadius: 12,
                        flexDirection: "column",
                        justifyContent: "flex-start",
                        alignItems: "flex-start",
                        gap: 8,
                        display: "inline-flex",
                      }}
                    >
                      <img
                        style={{ width: 120, height: 132, borderRadius: 8 }}
                        src="https://placehold.co/120x132"
                      />
                      <div
                        style={{
                          width: 120,
                          height: 36,
                          color: "black",
                          fontSize: 12,
                          fontFamily: "Noto Sans KR",
                          fontWeight: "900",
                          wordWrap: "break-word",
                        }}
                      >
                        공룡 뼈 조립 저 잘합니다. 도움이 필요하시면
                        말씀해주데오ㅛ
                      </div>
                    </div>
                    <div
                      data-속성-1="기본"
                      style={{
                        width: 136,
                        height: 207,
                        paddingTop: 12,
                        paddingBottom: 19,
                        paddingLeft: 8,
                        paddingRight: 8,
                        background: "white",
                        boxShadow: "0px 4px 20px rgba(32, 32, 32, 0.20)",
                        borderRadius: 12,
                        flexDirection: "column",
                        justifyContent: "flex-start",
                        alignItems: "flex-start",
                        gap: 8,
                        display: "inline-flex",
                      }}
                    >
                      <img
                        style={{ width: 120, height: 132, borderRadius: 8 }}
                        src="https://placehold.co/120x132"
                      />
                      <div
                        style={{
                          width: 120,
                          height: 36,
                          color: "black",
                          fontSize: 12,
                          fontFamily: "Noto Sans KR",
                          fontWeight: "900",
                          wordWrap: "break-word",
                        }}
                      >
                        제목글의 내용을 입력해주세요 여기에 제목글을 입력하면
                        자동으로 말줄임 됩니다
                      </div>
                    </div>
                    <div
                      data-속성-1="기본"
                      style={{
                        width: 136,
                        height: 207,
                        paddingTop: 12,
                        paddingBottom: 19,
                        paddingLeft: 8,
                        paddingRight: 8,
                        background: "white",
                        boxShadow: "0px 4px 20px rgba(32, 32, 32, 0.20)",
                        borderRadius: 12,
                        flexDirection: "column",
                        justifyContent: "flex-start",
                        alignItems: "flex-start",
                        gap: 8,
                        display: "inline-flex",
                      }}
                    >
                      <img
                        style={{ width: 120, height: 132, borderRadius: 8 }}
                        src="https://placehold.co/120x132"
                      />
                      <div
                        style={{
                          width: 120,
                          height: 36,
                          color: "black",
                          fontSize: 12,
                          fontFamily: "Noto Sans KR",
                          fontWeight: "900",
                          wordWrap: "break-word",
                        }}
                      >
                        제목글의 내용을 입력해주세요 여기에 제목글을 입력하면
                        자동으로 말줄임 됩니다
                      </div>
                    </div>
                  </div>
                </div>
                <div
                  style={{
                    width: 393,
                    height: 278,
                    left: 0,
                    top: 0,
                    position: "absolute",
                  }}
                />
              </div>
              <div
                style={{
                  alignSelf: "stretch",
                  height: 278,
                  position: "relative",
                }}
              >
                <div
                  style={{
                    width: 440,
                    left: 24,
                    top: 0,
                    position: "absolute",
                    flexDirection: "column",
                    justifyContent: "flex-start",
                    alignItems: "flex-start",
                    gap: 16,
                    display: "inline-flex",
                  }}
                >
                  <div
                    style={{
                      alignSelf: "stretch",
                      color: "black",
                      fontSize: 24,
                      fontFamily: "Noto Sans KR",
                      fontWeight: "900",
                      wordWrap: "break-word",
                    }}
                  >
                    실시간 인기 글
                  </div>
                  <div
                    style={{
                      alignSelf: "stretch",
                      justifyContent: "flex-start",
                      alignItems: "center",
                      gap: 12,
                      display: "inline-flex",
                    }}
                  >
                    <div
                      data-속성-1="기본"
                      style={{
                        width: 136,
                        height: 207,
                        paddingTop: 12,
                        paddingBottom: 19,
                        paddingLeft: 8,
                        paddingRight: 8,
                        background: "white",
                        boxShadow: "0px 4px 20px rgba(32, 32, 32, 0.20)",
                        borderRadius: 12,
                        flexDirection: "column",
                        justifyContent: "flex-start",
                        alignItems: "flex-start",
                        gap: 8,
                        display: "inline-flex",
                      }}
                    >
                      <img
                        style={{ width: 120, height: 132, borderRadius: 8 }}
                        src="https://placehold.co/120x132"
                      />
                      <div
                        style={{
                          width: 120,
                          height: 36,
                          color: "black",
                          fontSize: 12,
                          fontFamily: "Noto Sans KR",
                          fontWeight: "900",
                          wordWrap: "break-word",
                        }}
                      >
                        제목글의 내용을 입력해주세요 여기에 제목글을 입력하면
                        자동으로 말줄임 됩니다
                      </div>
                    </div>
                    <div
                      data-속성-1="기본"
                      style={{
                        width: 136,
                        height: 207,
                        paddingTop: 12,
                        paddingBottom: 19,
                        paddingLeft: 8,
                        paddingRight: 8,
                        background: "white",
                        boxShadow: "0px 4px 20px rgba(32, 32, 32, 0.20)",
                        borderRadius: 12,
                        flexDirection: "column",
                        justifyContent: "flex-start",
                        alignItems: "flex-start",
                        gap: 8,
                        display: "inline-flex",
                      }}
                    >
                      <img
                        style={{ width: 120, height: 132, borderRadius: 8 }}
                        src="https://placehold.co/120x132"
                      />
                      <div
                        style={{
                          width: 120,
                          height: 36,
                          color: "black",
                          fontSize: 12,
                          fontFamily: "Noto Sans KR",
                          fontWeight: "900",
                          wordWrap: "break-word",
                        }}
                      >
                        제목글의 내용을 입력해주세요 여기에 제목글을 입력하면
                        자동으로 말줄임 됩니다
                      </div>
                    </div>
                    <div
                      data-속성-1="기본"
                      style={{
                        width: 136,
                        height: 207,
                        paddingTop: 12,
                        paddingBottom: 19,
                        paddingLeft: 8,
                        paddingRight: 8,
                        background: "white",
                        boxShadow: "0px 4px 20px rgba(32, 32, 32, 0.20)",
                        borderRadius: 12,
                        flexDirection: "column",
                        justifyContent: "flex-start",
                        alignItems: "flex-start",
                        gap: 8,
                        display: "inline-flex",
                      }}
                    >
                      <img
                        style={{ width: 120, height: 132, borderRadius: 8 }}
                        src="https://placehold.co/120x132"
                      />
                      <div
                        style={{
                          width: 120,
                          height: 36,
                          color: "black",
                          fontSize: 12,
                          fontFamily: "Noto Sans KR",
                          fontWeight: "900",
                          wordWrap: "break-word",
                        }}
                      >
                        제목글의 내용을 입력해주세요 여기에 제목글을 입력하면
                        자동으로 말줄임 됩니다
                      </div>
                    </div>
                  </div>
                </div>
                <div
                  style={{
                    width: 393,
                    height: 278,
                    left: 0,
                    top: 0,
                    position: "absolute",
                  }}
                />
              </div>
            </div>
          </div>

          {/* <div
            style={{
              width: 393,
              paddingTop: 21,
              paddingBottom: 19,
              paddingLeft: 24,
              paddingRight: 24,
              left: 0,
              top: 0,
              position: "absolute",
              justifyContent: "center",
              alignItems: "center",
              gap: 154,
              display: "flex",
            }}
          >
            <div
              style={{
                flex: "1 1 0",
                height: 22,
                paddingTop: 1.5,
                justifyContent: "center",
                alignItems: "center",
                gap: 10,
                display: "flex",
              }}
            >
              <div
                style={{
                  textAlign: "center",
                  color: "#111111",
                  fontSize: 17,
                  fontFamily: "SF Pro",
                  fontWeight: "590",
                  lineHeight: 22,
                  wordWrap: "break-word",
                }}
              >
                9:41
              </div>
            </div>
            <div
              style={{
                flex: "1 1 0",
                height: 22,
                paddingTop: 1,
                paddingRight: 1,
                justifyContent: "center",
                alignItems: "center",
                gap: 7,
                display: "flex",
              }}
            >
              <div
                style={{ width: 19.2, height: 12.23, background: "#111111" }}
              />
              <div
                style={{ width: 17.14, height: 12.33, background: "#111111" }}
              />
              <div style={{ width: 27.33, height: 13, position: "relative" }}>
                <div
                  style={{
                    width: 25,
                    height: 13,
                    left: 0,
                    top: 0,
                    position: "absolute",
                    opacity: 0.35,
                    borderRadius: 4.3,
                    border: "1px #111111 solid",
                  }}
                />
                <div
                  style={{
                    width: 1.33,
                    height: 4.08,
                    left: 26,
                    top: 4.5,
                    position: "absolute",
                    opacity: 0.4,
                    background: "#111111",
                  }}
                />
                <div
                  style={{
                    width: 21,
                    height: 9,
                    left: 2,
                    top: 2,
                    position: "absolute",
                    background: "#111111",
                    borderRadius: 2.5,
                  }}
                />
              </div>
            </div>
          </div> */}
        </div>
      </header>
    </div>
  );
});

export default Home;
