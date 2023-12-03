import HistoryCards from './HistoryCards';

export default function HistoryBody() {

  return (
    <>
      <h1 className="font-bold text-4xl md:text-5xl lg:text-6xl mt-10 ml-20 text-indigo-500">File History</h1>
<section className="flex flex-row justify-center p-16 mb-36">
  <div className="w-full bg-gradient-to-r from-purple-300 to-pink-200 rounded-lg shadow-xl">
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 p-5">
      <HistoryCards />
    </div>
  </div>
</section>


    </>
  );
}
