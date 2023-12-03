"use client"

import { useState } from "react";

export default function HistoryCards() {
  const Cards = [
    { message: "aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf V aogmapinfwpaf V V aogmapinfwpaf VVaogmapin V V aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf V V aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf aogmapinfwpaf" },
    { message: "aogmapinfwpaf" },
    { message: "aogmapinfwpaf" },
    { message: "aogmapinfwpaf" },
  ];

  const [selectedCard, setSelectedCard] = useState(null);

  const handleCardClick = (index: any) => {
    setSelectedCard(index);
  };

  const handleCloseModal = () => {
    setSelectedCard(null);
  };

  return (
    <>
      {Cards.map((card, index) => (
        <div key={index} className="flex justify-center">
          <div
            className={`bg-white w-96 p-5 rounded-md cursor-pointer hover:scale-105 duration-200 ease-in-out overflow-hidden 
            ${selectedCard !== null ? "line-clamp-6" : "line-clamp-6"}`}
            onClick={() => handleCardClick(index)}
          >
            {card.message}
          </div>
        </div>
      ))}
      {selectedCard !== null && (
        
        <div className="fixed top-0 left-0 w-full h-full flex items-center justify-center bg-black bg-opacity-50 z-20">
          <div className="bg-white p-8 rounded-md w-[300px] sm:w-[400px] md:w-[700px] h-[400px] overflow-y-auto shadow-xl">
            <div className="flex justify-between">
              <button className=" bg-black text-white p-2 rounded-md hover:bg-zinc-700 mb-5 scale-90 md:scale-100" onClick={handleCloseModal}>Close</button><span className=" text-sm md:text-base">22/10/2023</span>
              <button className="  p-2 rounded-md hover:bg-gray-200 mb-5"><svg xmlns="http://www.w3.org/2000/svg" height="16" width="14" viewBox="0 0 448 512"><path d="M135.2 17.7L128 32H32C14.3 32 0 46.3 0 64S14.3 96 32 96H416c17.7 0 32-14.3 32-32s-14.3-32-32-32H320l-7.2-14.3C307.4 6.8 296.3 0 284.2 0H163.8c-12.1 0-23.2 6.8-28.6 17.7zM416 128H32L53.2 467c1.6 25.3 22.6 45 47.9 45H346.9c25.3 0 46.3-19.7 47.9-45L416 128z"/></svg></button>
            </div>
            <h1 className="font-bold text-3xl text-purple-400">New</h1>
            <p className="text-base">{Cards[selectedCard].message}</p>
            <h1 className="font-bold text-3xl text-black mt-5">Old</h1>
          </div>
        </div>
      )}
    </>
  );
}
