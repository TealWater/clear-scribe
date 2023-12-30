"use client"
import { useState } from "react";
import "../FAQ.css";

export default function Faq() {

  // array of questions and answers
  const datas = [
    {
      question: 'What is ClearScribe?',
      answer: 'Clear Scribe is an innovative app designed to simplify complex text, making it easy to understand and digest. It is your go-to solution for transforming intricate language into clear, concise, and easily understandable content.',
    },
    {
      question: 'How does it work?',
      answer: 'Clear Scribe utilizes algorithms and language processing techniques to break down complex sentences and words. The app then presents the information in a simplified and straightforward manner, making it more accessible to users.',
    },
    {
        question: 'Who can use it?',
        answer: 'From a student to a professional, everyone can use ClearScribe to simplify any text you want.',
      },
  ];
  // question at specific index to open it
  const [activeIndex, setActiveIndex] = useState(null);
  const handleToggleAccordion = (index : any) => {
    setActiveIndex(activeIndex === index ? null : index);
  };
  
  return (
      <div className="p-20">
        <div>
          <h1 className="accordion_1 text-6xl flex justify-center mb-10 font-semibold">FAQ&lsquo;s</h1>
        </div>
        <div className="accordion_2 space-y-2">
          {datas.map((item, index) =>  (
            <div key={index}>
                <button className="rounded" onClick={(  ) => handleToggleAccordion(index)}>
                <h3 style={{ fontWeight: 'bold', fontSize: '1.2rem' }}>{item.question}</h3> 
                { activeIndex === index ? ( 
                  <div className="p-2"><svg xmlns="http://www.w3.org/2000/svg" height="16" width="12" viewBox="0 0 384 512"><path d="M376.6 84.5c11.3-13.6 9.5-33.8-4.1-45.1s-33.8-9.5-45.1 4.1L192 206 56.6 43.5C45.3 29.9 25.1 28.1 11.5 39.4S-3.9 70.9 7.4 84.5L150.3 256 7.4 427.5c-11.3 13.6-9.5 33.8 4.1 45.1s33.8 9.5 45.1-4.1L192 306 327.4 468.5c11.3 13.6 31.5 15.4 45.1 4.1s15.4-31.5 4.1-45.1L233.7 256 376.6 84.5z"/></svg></div>
                 ): (
                  <div className="p-2"><svg xmlns="http://www.w3.org/2000/svg" height="16" width="12" viewBox="0 0 384 512"><path d="M169.4 470.6c12.5 12.5 32.8 12.5 45.3 0l160-160c12.5-12.5 12.5-32.8 0-45.3s-32.8-12.5-45.3 0L224 370.8 224 64c0-17.7-14.3-32-32-32s-32 14.3-32 32l0 306.7L54.6 265.4c-12.5-12.5-32.8-12.5-45.3 0s-12.5 32.8 0 45.3l160 160z"/></svg></div>
                 )}
                   </button>   
                  {activeIndex === index && (
                    <div className="answer-container">
                    <div className="answer-content">
                      <p>{item.answer}</p>
                    </div>
                  </div>
                )}
            </div>
          ))}
        </div>
      </div>
        );
}