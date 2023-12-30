"use client"

import { useState, useEffect } from "react";
interface Card {
  _id: string;
  createdAt: string;
  dateString: string;
  messageOld: string;
  messageNew: string;
}
let url = "http://localhost:8080";

export default function HistoryCards() {
  const [dataMessage, setDataMessage] = useState<Card[]>([]);
  const [selectedCard, setSelectedCard] = useState<Card | null>(null);

  // fetch from database and store it in dataMessage
  const fetchData = async () => {
    try {
      const res = await fetch(`${url}/history`);
      const data = await res.json();
      setDataMessage(data);
      console.log(data);
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  };

  // using id of the specific card selected sends a DELETE method and clears from the state 
  const deleteCard = async (id: string) => {
    try {
      const response = await fetch(`${url}/history?id=${id}`, {
        method: 'DELETE',
      });

      if (response.ok) {
        console.log('Resource deleted successfully');
        setDataMessage((prevData) => prevData.filter(card => card._id !== id));
      } else {
        console.error('Failed to delete resource');
      }
    } catch (error) {
      console.error('An error occurred:', error);
    }
  };

  useEffect(() => {
    fetchData();
  }, []);

  // depending on card clicked it opens up the modal with its specific info
  const handleCardClick = (card: any) => {
    setSelectedCard(card);
  };

  // changes card to null to close modal
  const handleCloseModal = () => {
    setSelectedCard(null);
  };

  if (dataMessage) {
    return (
      <>
        {dataMessage.map((card) => (
          <div key={card._id} className="flex justify-center">
            <div
              className={`bg-white w-96 p-5 rounded-md cursor-pointer hover:scale-105 duration-200 ease-in-out overflow-hidden 
            ${selectedCard !== null ? "line-clamp-6" : "line-clamp-6"}`}
              onClick={() => handleCardClick(card)}>
              {card.messageNew}
            </div>
          </div>
        ))}


        {selectedCard !== null && (
          <div className="fixed top-0 left-0 w-full h-full flex items-center justify-center bg-black bg-opacity-50 z-20">
            <div className="bg-white p-8 rounded-md w-[300px] sm:w-[400px] md:w-[700px] h-[400px] overflow-y-auto shadow-xl">
              <div className="flex md:justify-between ">
                <button
                  className="bg-black text-white p-2 rounded-md hover:bg-zinc-700 mb-5 scale-90 md:scale-100"
                  onClick={handleCloseModal}
                >
                  Close
                </button>
                <span className="text-sm text-center md:text-base">{selectedCard.dateString}</span>
                <button onClick={() => { handleCloseModal(); deleteCard(selectedCard._id); }} className="p-2 rounded-md hover:bg-gray-200 mb-5">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    height="16"
                    width="14"
                    viewBox="0 0 448 512"
                  >
                    <path d="M135.2 17.7L128 32H32C14.3 32 0 46.3 0 64S14.3 96 32 96H416c17.7 0 32-14.3 32-32s-14.3-32-32-32H320l-7.2-14.3C307.4 6.8 296.3 0 284.2 0H163.8c-12.1 0-23.2 6.8-28.6 17.7zM416 128H32L53.2 467c1.6 25.3 22.6 45 47.9 45H346.9c25.3 0 46.3-19.7 47.9-45L416 128z" />
                  </svg>
                </button>
              </div>
              <h1 className="font-bold text-3xl text-purple-400">New</h1>
              <div className="text-base">{selectedCard.messageNew}</div>
              <h1 className="font-bold text-3xl text-black mt-5">Old</h1>
              <p className="text-base">{selectedCard.messageOld}</p>
            </div>
          </div>
        )}
      </>
    );
  }
}
