import { useState } from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";

import MainLayout from "./layout/mainLayout";
import BodyCard from "./components/body-card";
import NurseCall from "./pages/NurseCall";
import Logs from "./pages/Logs";
import Setting from "./pages/Setting";
function App() {
  return (
    <div className="bg-slate-900 h-screen">
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<NurseCall />} />
          <Route path="/Home" element={<NurseCall />} />
          <Route path="/Logs" element={<Logs />} />
          <Route path="/Setting" element={<Setting />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
