import Sidebar from "./components/Sidebar.jsx";
import Navbar from "./components/Navbar.jsx";
import NurseCallSimulator from "./components/NurseCall.jsx";
import CardSlider from "./components/CardSlider.jsx";
import NurseCallPage from "./pages/Nursecall.jsx";

function App() {
  const Components = [<NurseCallSimulator />, <Sidebar preview={false} />];
  return (
    <div className="bg-slate-900 h-screen  ">
      {/* <div className="bg-black h-screen">
        <div className="h-30 ">
          <Navbar />
        </div>
        <div className="grid place-items-center p-2 mt-25  ">
          <div className="flex items-center">
            <div className=" mx-2 my-3">
              <CardSlider slides={Components} />
            </div>
          </div>
        </div>
      </div> */}
      <NurseCallPage />
    </div>
  );
}

export default App;
