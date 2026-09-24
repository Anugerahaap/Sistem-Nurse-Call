import React, { useState } from "react";
import { FaChevronLeft, FaChevronRight } from "react-icons/fa6";

export default function CardSlider({ slides = [] }) {
  const [currentIndex, setCurrentIndex] = useState(0);
  const prevSlide = () => {
    const isFirstSlide = currentIndex === 0;
    const newIndex = isFirstSlide ? slides.length - 1 : currentIndex - 1;
    setCurrentIndex(newIndex);
  };
  const nextSlide = () => {
    const isLastSlide = currentIndex === slides.length - 1;
    const newIndex = isLastSlide ? 0 : currentIndex + 1;
    setCurrentIndex(newIndex);
  };

  return (
    <div className="w-[482px] max-h-[572px] m-auto py-14 px-4 relative border border-slate-400 bg-gray-300 shadow-sl rounded-2xl group overflow-hidden">
      <div className="hidden group-hover:block absolute top-[50%] -translate-y-[-50%] left-5 text-2xl rounded-full bg-black/20 text-white cursor-pointer p-2">
        <FaChevronLeft size={30} onClick={prevSlide} />
      </div>
      <div className="hidden group-hover:block absolute top-[50%] -translate-y-[-50%] right-5 text-2xl rounded-full  bg-black/20 text-white cursor-pointer p-2 ml-4">
        <FaChevronRight size={30} onClick={nextSlide} />
      </div>
      <div className="w-full h-full rounded-2xl bg-center bg-cover duration-500 ">
        {slides[currentIndex]}
      </div>
    </div>
  );
}
