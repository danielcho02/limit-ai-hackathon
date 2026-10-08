import { useState } from "react";

export default function Carousel({ data }: { data: any }) {
  const [currentIndex, setCurrentIndex] = useState(0);
  const src = require("@/assets/images/images.png");
  return (
    <div className="carousel-container">
      {data?.length > 0 && (
        // <img src={data[currentIndex].src} alt={data[currentIndex].alt || ""} />
        <img src={src} alt={data[currentIndex].alt || ""} />
      )}
      {/* 화살표 버튼을 추가하여 좌우로 슬라이드를 이동할 수 있게 합니다 */}
      <button
        onClick={() => currentIndex > 0 && setCurrentIndex(currentIndex - 1)}
      >
        Prev
      </button>
      <button
        disabled={currentIndex === data.length - 1}
        onClick={() => setCurrentIndex(currentIndex + 1)}
      >
        Next
      </button>
    </div>
  );
}
