"use client"


import { useState, useEffect } from "react";
interface Card {
  id: number;
  createdAt: string;
  messageOld: string;
  messageNew: string;
}

export default function HistoryCards() {
  //const [dataMessage, setDataMessage] = useState([]);
  const [selectedCard, setSelectedCard] = useState<Card | null>(null);

  const dataMessage = [
    {
        "id": 0,
        "createdAt": "2023-12-03 14:28:07.809752 -0500 EST m=+11.133889846",
        "messageOld": "I like taking a stroll down mempry lane",
        "messageNew": "I like taking a walk down memory lane"
    },
    {
        "id": 1,
        "createdAt": "2023-12-03 14:28:07.809758 -0500 EST m=+11.133895769",
        "messageOld": "All humans have gone through a period of gestation for nine months",
        "messageNew": "All humans have gone through a period of development for nine months"
    },
    {
        "id": 2,
        "createdAt": "2023-12-03 14:28:07.809759 -0500 EST m=+11.133896818",
        "messageOld": "I have no quarrel with Cammalot",
        "messageNew": "I have no problem with Cammalot"
    },
    {
        "id": 3,
        "createdAt": "2023-12-03 14:28:07.80976 -0500 EST m=+11.133897519",
        "messageOld": "Do you have any more queries?",
        "messageNew": "Do you have any more questions?"
    },
    {
        "id": 4,
        "createdAt": "2023-12-03 14:28:07.80976 -0500 EST m=+11.133898268",
        "messageOld": "My classroom was adjacent to the library.",
        "messageNew": "My classroom was next to the library."
    },
    {
        "id": 5,
        "createdAt": "2023-12-03 14:28:07.809761 -0500 EST m=+11.133898878",
        "messageOld": "The child has a inqusistive look.",
        "messageNew": "The child has a pensive look."
    }
]


  /* useEffect(() => {
    const fetchData = async () => {
      try {
        const res = await fetch("http://localhost:8080/history");
        const data = await res.json();
        setDataMessage(data);
        console.log(data);
      } catch (error) {
        console.error("Error fetching data:", error);
      }
    };

    fetchData(); 
  }, []); */

  const handleCardClick = (card: any) => {
    setSelectedCard(card);

  };

  const handleCloseModal = () => {
    setSelectedCard(null);
  };

  return (
    <>

      {dataMessage.map((card, index) => (

        <div key={index} className="flex justify-center">
          <div
            className={`bg-white w-96 p-5 rounded-md cursor-pointer hover:scale-105 duration-200 ease-in-out overflow-hidden 
            ${selectedCard !== null ? "line-clamp-6" : "line-clamp-6"}`}

            onClick={() => handleCardClick(card)}
          >
            {card.messageNew}
          </div>
        </div>
      ))}

      
{selectedCard !== null && (
  <div className="fixed top-0 left-0 w-full h-full flex items-center justify-center bg-black bg-opacity-50 z-20">
    <div className="bg-white p-8 rounded-md w-[300px] sm:w-[400px] md:w-[700px] h-[400px] overflow-y-auto shadow-xl">
      <div className="flex justify-between">
        <button
          className="bg-black text-white p-2 rounded-md hover:bg-zinc-700 mb-5 scale-90 md:scale-100"
          onClick={handleCloseModal}
        >
          Close
        </button>
        <span className="text-sm md:text-base">22/10/2023</span>
        <button className="p-2 rounded-md hover:bg-gray-200 mb-5">
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
