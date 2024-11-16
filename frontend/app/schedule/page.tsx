import CalendarComponent from "../components/Calendar";
import { Navbar } from "../components/Navbar";

export default function Home() {
    return (
        <div>
            <Navbar />
            <h1 className="text-6xl mt-10 text-center">Schedule</h1>
            <CalendarComponent />
        </div>
    );
}