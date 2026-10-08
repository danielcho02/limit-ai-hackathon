import { useMemo, useState, type ReactNode } from "react";
import logo from "./logo.png";

type Page = "home" | "market" | "chat" | "my";
type IconName =
  | "home"
  | "market"
  | "chat"
  | "user"
  | "search"
  | "bell"
  | "heart"
  | "message"
  | "arrow"
  | "plus"
  | "send"
  | "close"
  | "spark"
  | "chevron"
  | "edit";

const photos = {
  craft:
    "https://images.unsplash.com/photo-1506806732259-39c2d0268443?auto=format&fit=crop&w=1200&q=85",
  thread:
    "https://images.unsplash.com/photo-1522065893269-6fd20f6d7438?auto=format&fit=crop&w=900&q=85",
  class:
    "https://images.unsplash.com/photo-1560831340-b9679dc9e9f0?auto=format&fit=crop&w=900&q=85",
  ceramic:
    "https://images.unsplash.com/photo-1480355781839-51097c7d4f9f?auto=format&fit=crop&w=900&q=85",
  logo: "./logo.png",
};

const feed = [
  {
    id: 1,
    type: "재능 낭비",
    user: "반듯한삐뚤이",
    meta: "서울 마포구 · 12분 전",
    title: "영수증만 보면 시가 떠오릅니다",
    body: "버리기 전 영수증을 보내주세요. 오늘의 소비를 4행시로 만들어 드려요.",
    image: photos.craft,
    tags: ["#쓸데없는", "#글쓰기", "#3분완성"],
    likes: 128,
    comments: 23,
    rank: "오늘의 낭비 1위",
  },
  {
    id: 2,
    type: "재능 요청",
    user: "고민많은감자",
    meta: "온라인 · 34분 전",
    title: "면접 전, 제 목소리에 자신감을 넣어주세요",
    body: "첫 면접을 앞두고 있어요. 딱 10분만 제 이야기를 듣고 용기를 빌려주실 분!",
    image: photos.thread,
    tags: ["#상담", "#취준", "#목소리"],
    likes: 74,
    comments: 16,
    rank: "매칭 급상승",
  },
  {
    id: 3,
    type: "재능 낭비",
    user: "종이접는악마",
    meta: "부산 수영구 · 1시간 전",
    title: "안 쓰는 종이로 쓸데없이 귀여운 것 만들기",
    body: "종이 한 장과 20분이면 충분해요. 지역아동센터 단체 수업도 환영합니다.",
    image: photos.class,
    tags: ["#문화예술", "#교육", "#기관환영"],
    likes: 96,
    comments: 31,
    rank: "기관 추천",
  },
  {
    id: 4,
    type: "재능 낭비",
    user: "그릇된재능",
    meta: "전국 택배 · 2시간 전",
    title: "깨진 마음 말고, 깨진 그릇을 이어드려요",
    body: "작은 흠집이 생긴 도자기를 새로운 무늬로 고쳐드리는 소소한 재능입니다.",
    image: photos.ceramic,
    tags: ["#환경", "#공예", "#업사이클"],
    likes: 211,
    comments: 42,
    rank: "주간 베스트",
  },
];

const chats = [
  {
    name: "고현준",
    message: "그럼 수요일 3시에 뵐까요?",
    time: "방금",
    unread: 2,
    color: "bg-[#E7B47A]",
  },
  {
    name: "양동균",
    message: "안녕하세요! 제게도 도와주실 수 있나요?",
    time: "12분",
    unread: 1,
    color: "bg-[#D33024]",
  },
  {
    name: "종이접는악마",
    message: "사진 확인했습니다. 너무 좋아요!",
    time: "어제",
    unread: 0,
    color: "bg-[#587A63]",
  },
  {
    name: "마음정리함",
    message: "따뜻한 후기 정말 감사해요.",
    time: "월",
    unread: 0,
    color: "bg-[#9186AD]",
  },
];

function Icon({
  name,
  size = 22,
  filled = false,
}: {
  name: IconName;
  size?: number;
  filled?: boolean;
}) {
  const paths: Record<IconName, ReactNode> = {
    home: (
      <>
        <path d="M3 10.8 12 3l9 7.8" />
        <path d="M5.5 9.5V21h13V9.5M9 21v-7h6v7" />
      </>
    ),
    market: (
      <>
        <path d="M4 10v10h16V10M3 4h18l-1.5 6H4.5L3 4Z" />
        <path d="M8 10V4m8 6V4M8 15h8" />
      </>
    ),
    chat: <path d="M4 4h16v12H9l-5 4V4Z" />,
    user: (
      <>
        <circle cx="12" cy="8" r="4" />
        <path d="M4 21c.6-4 3.2-6 8-6s7.4 2 8 6" />
      </>
    ),
    search: (
      <>
        <circle cx="10.5" cy="10.5" r="6.5" />
        <path d="m16 16 5 5" />
      </>
    ),
    bell: (
      <>
        <path d="M6 17h12l-1.5-2.5V10a4.5 4.5 0 0 0-9 0v4.5L6 17Z" />
        <path d="M10 20h4" />
      </>
    ),
    heart: (
      <path d="M20.8 5.7c-2.1-2.1-5.5-2.1-7.6 0L12 6.9l-1.2-1.2a5.4 5.4 0 0 0-7.6 7.6L12 22l8.8-8.7a5.4 5.4 0 0 0 0-7.6Z" />
    ),
    message: <path d="M4 5h16v12H9l-5 4V5Z" />,
    arrow: (
      <>
        <path d="m5 12 14 0" />
        <path d="m14 7 5 5-5 5" />
      </>
    ),
    plus: (
      <>
        <path d="M12 5v14M5 12h14" />
      </>
    ),
    send: (
      <>
        <path d="m3 11 18-8-7 18-3-7-8-3Z" />
        <path d="m11 14 4-4" />
      </>
    ),
    close: (
      <>
        <path d="m5 5 14 14M19 5 5 19" />
      </>
    ),
    spark: (
      <>
        <path d="m12 2 1.5 5.5L19 9l-5.5 1.5L12 16l-1.5-5.5L5 9l5.5-1.5L12 2Z" />
        <path d="m19 16 .7 2.3L22 19l-2.3.7L19 22l-.7-2.3L16 19l2.3-.7L19 16Z" />
      </>
    ),
    chevron: <path d="m9 6 6 6-6 6" />,
    edit: (
      <>
        <path d="m4 20 4.5-1 11-11-3.5-3.5-11 11L4 20Z" />
        <path d="m14.5 6 3.5 3.5" />
      </>
    ),
  };
  return (
    <svg
      aria-hidden="true"
      width={size}
      height={size}
      viewBox="0 0 24 24"
      fill={filled ? "currentColor" : "none"}
      stroke="currentColor"
      strokeWidth="1.8"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      {paths[name]}
    </svg>
  );
}

function Logo({ light = false }: { light?: boolean }) {
  return (
    <img src={logo} />
    // <div className={`logo ${light ? "text-white" : "text-[#D33024]"}`}>
    //   악마의 재능<span>쓸모없어 더 특별한</span>
    // </div>
  );
}

function NavItem({
  active,
  icon,
  label,
  onClick,
}: {
  active: boolean;
  icon: IconName;
  label: string;
  onClick: () => void;
}) {
  return (
    <button className={`nav-item ${active ? "active" : ""}`} onClick={onClick}>
      <Icon name={icon} size={21} filled={active && icon === "home"} />
      <span>{label}</span>
    </button>
  );
}

function Sidebar({
  page,
  setPage,
  onCreate,
}: {
  page: Page;
  setPage: (p: Page) => void;
  onCreate: () => void;
}) {
  return (
    <aside className="sidebar">
      <Logo />
      <nav className="side-nav">
        <NavItem
          active={page === "home"}
          icon="home"
          label="홈"
          onClick={() => setPage("home")}
        />
        <NavItem
          active={page === "market"}
          icon="market"
          label="재능 마켓"
          onClick={() => setPage("market")}
        />
        <NavItem
          active={page === "chat"}
          icon="chat"
          label="채팅"
          onClick={() => setPage("chat")}
        />
        <NavItem
          active={page === "my"}
          icon="user"
          label="마이페이지"
          onClick={() => setPage("my")}
        />
      </nav>
      <button className="waste-button" onClick={onCreate}>
        <Icon name="plus" size={20} />
        재능 낭비하기
      </button>
      <div className="side-profile">
        <div className="avatar avatar-dark">낭</div>
        <div>
          <strong>낭비의신</strong>
          <span>Lv. 7 · 하급 악마</span>
        </div>
        <Icon name="chevron" size={17} />
      </div>
    </aside>
  );
}

function Topbar({ title, subtitle }: { title: string; subtitle?: string }) {
  return (
    <header className="topbar">
      <div>
        <span className="eyebrow">{subtitle || "악마의 재능"}</span>
        <h1>{title}</h1>
      </div>
      <div className="top-actions">
        <button aria-label="검색">
          <Icon name="search" />
        </button>
        <button aria-label="알림" className="notification">
          <Icon name="bell" />
          <i />
        </button>
      </div>
    </header>
  );
}

function Home({
  onMarket,
  onCreate,
}: {
  onMarket: () => void;
  onCreate: () => void;
}) {
  const [category, setCategory] = useState("전체");
  const [liked, setLiked] = useState<number[]>([4]);
  const categories = ["전체", "문화예술", "환경", "교육", "상담", "쓸데없는"];
  const visible =
    category === "전체"
      ? feed
      : feed.filter((item) => item.tags.some((tag) => tag.includes(category)));

  const toggleLike = (id: number) =>
    setLiked((current) =>
      current.includes(id)
        ? current.filter((item) => item !== id)
        : [...current, id],
    );

  return (
    <main className="page">
      <Topbar
        title="오늘은 뭘 낭비해볼까요?"
        subtitle="GOOD WASTE, GREAT TALENT"
      />
      <section className="hero">
        <div className="hero-copy">
          <span className="hero-kicker">
            <Icon name="spark" size={16} />
            이번 주 1,284개의 재능이 낭비됐어요
          </span>
          <h2>
            사소한 재능이
            <br />
            누군가에겐 <em>사건</em>이 된다.
          </h2>
          <p>잘할 필요는 없어요. 당신만 할 수 있으면 되니까.</p>
          <div className="hero-buttons">
            <button className="primary-button" onClick={onCreate}>
              내 재능 낭비하기 <Icon name="arrow" size={18} />
            </button>
            <button className="text-button" onClick={onMarket}>
              도움 구걸하러 가기
            </button>
          </div>
        </div>
      </section>

      <section className="section-block">
        <div className="section-heading">
          <div>
            <span className="eyebrow">LIVE CATEGORY</span>
            <h2>어떤 재능을 찾으세요?</h2>
          </div>
          <button className="text-link" onClick={onMarket}>
            전체 보기 <Icon name="arrow" size={16} />
          </button>
        </div>
        <div className="category-row">
          {categories.map((item, index) => (
            <button
              key={item}
              className={`category-chip ${category === item ? "selected" : ""}`}
              onClick={() => setCategory(item)}
            >
              <span>{["ALL", "ART", "ECO", "EDU", "TALK", "ODD"][index]}</span>
              {item}
            </button>
          ))}
        </div>
      </section>

      <section className="section-block feed-section">
        <div className="section-heading">
          <div>
            <span className="eyebrow">TRENDING NOW</span>
            <h2>지금 불타는 재능</h2>
          </div>
          <span className="live-label">
            <i /> 실시간 업데이트
          </span>
        </div>
        <div className="feed-grid">
          {visible.map((item, index) => (
            <article
              className={`talent-card ${index === 0 ? "featured" : ""}`}
              key={item.id}
            >
              <div className="card-image">
                <img src={item.image} alt="" />
                <span
                  className={`type-badge ${item.type === "재능 요청" ? "request" : ""}`}
                >
                  {item.type}
                </span>
                <span className="rank-badge">{item.rank}</span>
              </div>
              <div className="card-content">
                <div className="author-line">
                  <div className="avatar">{item.user.slice(0, 1)}</div>
                  <div>
                    <strong>{item.user}</strong>
                    <span>{item.meta}</span>
                  </div>
                </div>
                <h3>{item.title}</h3>
                <p>{item.body}</p>
                <div className="tags">
                  {item.tags.map((tag) => (
                    <span key={tag}>{tag}</span>
                  ))}
                </div>
                <div className="card-footer">
                  <button
                    className={liked.includes(item.id) ? "liked" : ""}
                    onClick={() => toggleLike(item.id)}
                  >
                    <Icon
                      name="heart"
                      size={19}
                      filled={liked.includes(item.id)}
                    />{" "}
                    {item.likes + (liked.includes(item.id) ? 1 : 0)}
                  </button>
                  <button>
                    <Icon name="message" size={19} /> {item.comments}
                  </button>
                  <button className="match-button">
                    매칭 제안 <Icon name="arrow" size={16} />
                  </button>
                </div>
              </div>
            </article>
          ))}
        </div>
      </section>
    </main>
  );
}

function Market({ onCreate }: { onCreate: () => void }) {
  const [tab, setTab] = useState<"offer" | "request">("offer");
  const [query, setQuery] = useState("");
  const items = useMemo(
    () =>
      feed.filter(
        (item) =>
          item.type === (tab === "offer" ? "재능 낭비" : "재능 요청") &&
          item.title.includes(query),
      ),
    [tab, query],
  );
  return (
    <main className="page">
      <Topbar title="재능 마켓" subtitle="TALENT FLEA MARKET" />
      <div className="market-toolbar">
        <div className="segmented">
          <button
            className={tab === "offer" ? "selected" : ""}
            onClick={() => setTab("offer")}
          >
            재능 줍기{" "}
            <span>{feed.filter((i) => i.type === "재능 낭비").length}</span>
          </button>
          <button
            className={tab === "request" ? "selected" : ""}
            onClick={() => setTab("request")}
          >
            재능 요청{" "}
            <span>{feed.filter((i) => i.type === "재능 요청").length}</span>
          </button>
        </div>
        <label className="search-box">
          <Icon name="search" size={20} />
          <input
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            placeholder="재능, 태그, 지역을 검색하세요"
          />
        </label>
      </div>
      <div className="market-intro">
        <div>
          <span className="eyebrow">
            {tab === "offer" ? "WASTE IT" : "ASK FOR IT"}
          </span>
          <h2>
            {tab === "offer"
              ? "누군가 버린 재능을 주워보세요."
              : "도움이 필요하다면 당당하게 구걸하세요."}
          </h2>
        </div>
        <button className="primary-button" onClick={onCreate}>
          <Icon name="plus" size={18} /> 글쓰기
        </button>
      </div>
      <div className="list-grid">
        {items.map((item) => (
          <article className="list-card" key={item.id}>
            <img src={item.image} alt="" />
            <div className="list-copy">
              <span className="type-text">{item.type}</span>
              <h3>{item.title}</h3>
              <p>{item.body}</p>
              <div className="tags">
                {item.tags.map((tag) => (
                  <span key={tag}>{tag}</span>
                ))}
              </div>
              <div className="list-meta">
                <strong>{item.user}</strong>
                <span>
                  <Icon name="heart" size={15} /> {item.likes}
                </span>
              </div>
            </div>
          </article>
        ))}
        {!items.length && (
          <div className="empty">
            찾는 재능이 없어요. 다른 단어로 낭비해보세요.
          </div>
        )}
      </div>
    </main>
  );
}

function Chat() {
  const [active, setActive] = useState(chats[0]);
  const [message, setMessage] = useState("");
  const [messages, setMessages] = useState([
    "안녕하세요! 물병 세우기 재능 보고 연락드렸어요.",
    "무슨 요일에 몇 시에 시간 되시나요?",
    "저는 월요일 4시 이후, 화·수는 모든 시간대 가능합니다!",
    "그럼 수요일 3시에 뵐까요?",
  ]);
  const send = () => {
    if (!message.trim()) return;
    setMessages((current) => [...current, message]);
    setMessage("");
  };
  return (
    <main className="page chat-page">
      <Topbar title="채팅" subtitle="YOUR MATCHES" />
      <div className="chat-layout">
        <section className="chat-list">
          <div className="chat-list-title">
            <strong>메시지</strong>
            <span>읽지 않음 3</span>
          </div>
          {chats.map((chat) => (
            <button
              key={chat.name}
              className={`chat-row ${active.name === chat.name ? "selected" : ""}`}
              onClick={() => setActive(chat)}
            >
              <div className={`avatar ${chat.color}`}>{chat.name[0]}</div>
              <div>
                <strong>{chat.name}</strong>
                <p>{chat.message}</p>
              </div>
              <div className="chat-status">
                <span>{chat.time}</span>
                {chat.unread > 0 && <i>{chat.unread}</i>}
              </div>
            </button>
          ))}
        </section>
        <section className="conversation">
          <header className="conversation-header">
            <div className={`avatar ${active.color}`}>{active.name[0]}</div>
            <div>
              <strong>{active.name}</strong>
              <span>
                <i /> 매칭 진행 중
              </span>
            </div>
          </header>
          <div className="safety-note">
            현금 거래나 대가 요구는 영구 정지 대상입니다. 안전한 재능 낭비를
            부탁드려요.
          </div>
          <div className="messages">
            <div className="date-divider">오늘</div>
            {messages.map((text, index) => (
              <div
                key={`${text}-${index}`}
                className={`bubble ${index % 3 === 2 || index === messages.length - 1 ? "mine" : ""}`}
              >
                {text}
                <span>{index === messages.length - 1 ? "오후 3:42" : ""}</span>
              </div>
            ))}
          </div>
          <div className="message-input">
            <button aria-label="파일 첨부">
              <Icon name="plus" size={20} />
            </button>
            <input
              value={message}
              onChange={(e) => setMessage(e.target.value)}
              onKeyDown={(e) => e.key === "Enter" && send()}
              placeholder="메시지를 입력하세요"
            />
            <button className="send-button" onClick={send} aria-label="전송">
              <Icon name="send" size={19} />
            </button>
          </div>
        </section>
      </div>
    </main>
  );
}

function MyPage() {
  return (
    <main className="page">
      <Topbar title="마이페이지" subtitle="MY DEVILISH PROFILE" />
      <section className="profile-hero">
        <div className="profile-info">
          <div className="big-avatar">
            낭<span>LV.7</span>
          </div>
          <div>
            <span className="eyebrow">하급 악마 · 다음 레벨까지 320P</span>
            <h2>낭비의신</h2>
            <p>쓸데없지만 꽤 진심인 재능 수집가입니다.</p>
            <div className="follow">
              <strong>
                128 <span>팔로워</span>
              </strong>
              <strong>
                64 <span>팔로잉</span>
              </strong>
            </div>
          </div>
          <button className="outline-button">
            <Icon name="edit" size={17} /> 프로필 수정
          </button>
        </div>
        <div className="level-bar">
          <span style={{ width: "68%" }} />
          <div>
            <b>680P</b>
            <em>1,000P</em>
          </div>
        </div>
      </section>
      <section className="stats-grid">
        <div className="stat-card red">
          <span>이번 주 천사링</span>
          <strong>17위</strong>
          <p>지난주보다 12계단 상승</p>
        </div>
        <div className="stat-card">
          <span>완료한 매칭</span>
          <strong>24</strong>
          <p>도망가지 않고 완료했어요</p>
        </div>
        <div className="stat-card">
          <span>매너 온도</span>
          <strong>48.7°</strong>
          <p>상위 8%의 따뜻함</p>
        </div>
      </section>
      <section className="achievements">
        <div className="section-heading">
          <div>
            <span className="eyebrow">MY ACHIEVEMENTS</span>
            <h2>하찮지만 빛나는 업적</h2>
          </div>
          <button className="text-link">
            전체 보기 <Icon name="arrow" size={16} />
          </button>
        </div>
        <div className="badge-row">
          {[
            ["첫 낭비", "재능 첫 등록 완료", "01"],
            ["이웃 사랑", "매칭 10회 완료", "10"],
            ["노쇼 제로", "약속 20회 연속 준수", "20"],
            ["반응 부자", "좋아요 100개 달성", "♥"],
          ].map(([title, text, mark]) => (
            <div className="achievement" key={title}>
              <div>{mark}</div>
              <strong>{title}</strong>
              <span>{text}</span>
            </div>
          ))}
        </div>
      </section>
      <section className="certificate">
        <div>
          <span className="eyebrow">WEEKLY REPORT</span>
          <h2>
            이번 주 당신의 재능은
            <br />
            <em>4명</em>을 웃게 했어요.
          </h2>
          <p>사소한 재능도 나누면 쓸모가 생겨요. 다음 낭비도 기대할게요.</p>
        </div>
        <div className="seal">
          GOOD
          <br />
          WASTE<span>악마의 재능 인증</span>
        </div>
      </section>
    </main>
  );
}

function CreateModal({ onClose }: { onClose: () => void }) {
  const [mode, setMode] = useState<"offer" | "request">("offer");
  const [done, setDone] = useState(false);
  if (done)
    return (
      <div className="modal-backdrop">
        <div className="success-modal">
          <div className="success-mark">
            <Icon name="spark" size={30} />
          </div>
          <span className="eyebrow">NICE WASTE!</span>
          <h2>재능이 성공적으로 낭비됐어요.</h2>
          <p>당신의 사소한 재능이 필요한 사람에게 곧 닿을 거예요.</p>
          <button className="primary-button" onClick={onClose}>
            피드로 돌아가기
          </button>
        </div>
      </div>
    );
  return (
    <div className="modal-backdrop" onMouseDown={onClose}>
      <div className="create-modal" onMouseDown={(e) => e.stopPropagation()}>
        <div className="modal-header">
          <div>
            <span className="eyebrow">NEW TALENT</span>
            <h2>무엇을 낭비할까요?</h2>
          </div>
          <button onClick={onClose} aria-label="닫기">
            <Icon name="close" />
          </button>
        </div>
        <div className="segmented full">
          <button
            className={mode === "offer" ? "selected" : ""}
            onClick={() => setMode("offer")}
          >
            재능 주기
          </button>
          <button
            className={mode === "request" ? "selected" : ""}
            onClick={() => setMode("request")}
          >
            재능 요청
          </button>
        </div>
        <label className="form-field">
          <span>제목</span>
          <input
            placeholder={
              mode === "offer"
                ? "낭비할 재능을 한 줄로 소개해주세요"
                : "어떤 도움이 필요한가요?"
            }
          />
        </label>
        <label className="form-field">
          <span>사연</span>
          <textarea
            placeholder="사소할수록 좋아요. 이야기를 들려주세요."
            rows={5}
          />
        </label>
        <div className="upload-row">
          <button>
            <Icon name="plus" /> 사진 또는 영상 추가
          </button>
          <span>최대 3개</span>
        </div>
        <label className="form-field">
          <span>
            태그 <em>3개 이상</em>
          </span>
          <input placeholder="#쓸데없는 #예술 #서울" />
        </label>
        <button
          className="primary-button modal-submit"
          onClick={() => setDone(true)}
        >
          {mode === "offer" ? "세상에 낭비하기" : "도움 구걸하기"}{" "}
          <Icon name="arrow" size={18} />
        </button>
      </div>
    </div>
  );
}

export default function App() {
  const [page, setPage] = useState<Page>("home");
  const [createOpen, setCreateOpen] = useState(false);
  const titles: Record<Page, string> = {
    home: "홈",
    market: "마켓",
    chat: "채팅",
    my: "MY",
  };
  return (
    <div className="app-shell">
      <Sidebar
        page={page}
        setPage={setPage}
        onCreate={() => setCreateOpen(true)}
      />
      <div className="content-shell">
        {page === "home" && (
          <Home
            onMarket={() => setPage("market")}
            onCreate={() => setCreateOpen(true)}
          />
        )}
        {page === "market" && <Market onCreate={() => setCreateOpen(true)} />}
        {page === "chat" && <Chat />}
        {page === "my" && <MyPage />}
      </div>
      <nav className="mobile-nav">
        {(["home", "market", "chat", "my"] as Page[]).map((item) => (
          <NavItem
            key={item}
            active={page === item}
            icon={item === "my" ? "user" : item}
            label={titles[item]}
            onClick={() => setPage(item)}
          />
        ))}
      </nav>
      <button
        className="mobile-create"
        onClick={() => setCreateOpen(true)}
        aria-label="재능 등록"
      >
        <Icon name="plus" />
      </button>
      {createOpen && <CreateModal onClose={() => setCreateOpen(false)} />}
    </div>
  );
}
