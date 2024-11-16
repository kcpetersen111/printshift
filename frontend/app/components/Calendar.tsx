"use client";

import React, { useEffect, useState } from "react";
import { Calendar, momentLocalizer, Event } from "react-big-calendar";
import moment from "moment";
import "react-big-calendar/lib/css/react-big-calendar.css"; // Base styles
import "../styles/react-big-calendar.css"; // Custom styles
import { SchedulePrinter } from "./SchedulePrinter";

const localizer = momentLocalizer(moment);

interface CustomEvent extends Event {
  title: string;
  start: Date;
  end: Date;
}

type Times = {
    startTime: Date,
    endTime: Date,
}

const CalendarComponent: React.FC = () => {
  const [events, setEvents] = useState<CustomEvent[]>([
    // {
    //   title: "Team Meeting",
    //   start: new Date(),
    //   end: new Date(new Date().setHours(new Date().getHours() + 1)),
    // },
    // {
    //   title: "Lunch Break",
    //   start: new Date(new Date().setHours(12)),
    //   end: new Date(new Date().setHours(13)),
    // },
  ]);
  const dummyTime: Times = {
    startTime: new Date(),
    endTime: new Date()
  }
  const [title, setTitle] = useState("");
  const [isOpen, setIsOpen] = useState(false);
  const [times, setTimes] = useState<Times>(dummyTime)

    const handleChangePrinter = (startTime: Date, endTime: Date) => {
        setTimes((prev) => ({ ...prev, startTime, endTime}));
    };

  const handleSelectSlot = (slotInfo: { start: Date; end: Date }) => {
    setIsOpen(true);

    handleChangePrinter(slotInfo.start, slotInfo.end);

    // fetch("http://localhost:3410/protected/createAvailablePrinterTime", {
    //     method: "POST",
    //     body: JSON.stringify(printerRequest)
    // });
    // fetch("http://localhost:3410/protected/createAvailableClassTime", {
    //     method: "POST",
    //     body: JSON.stringify(classRequest)
    // });
    // if (title) {
    //   setEvents([
    //     ...events,
    //     {
    //       title,
    //       start: slotInfo.start,
    //       end: slotInfo.end,
    //     },
    //   ]);
    // }
  };

  useEffect(() => {
    setEvents([
        ...events,
        {
            title: title,
            start: times.startTime,
            end: times.endTime
        }
    ]);
  }, [title])

  return (
    <div className="container mx-auto p-6">
      <div className="bg-white shadow-lg rounded-lg p-4">
        <Calendar
          localizer={localizer}
          events={events}
          startAccessor="start"
          endAccessor="end"
          selectable
          onSelectSlot={handleSelectSlot}
          style={{ height: "80vh" }}
          className="custom-calendar dark:text-gray-800"
        />
        <SchedulePrinter isOpen={isOpen} setIsOpen={setIsOpen} setTitle={setTitle} />
      </div>
    </div>
  );
};

export default CalendarComponent;
