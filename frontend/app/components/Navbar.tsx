"use client";
import Link from "next/link";

export function Navbar() {
    return (
        <div style={{ backgroundColor: "#003058" }} className="text-white p-4 w-full">
            <div className="container mx-auto flex justify-between items-center">
                <div className="container mx-auto flex justify-between items-center">
                    <h1 className="text-xl font-bold">
                        <Link href="/home">
                        <img 
                            src="printshift-logo.png" 
                            alt="Printshift logo"
                            width="27%"
                        />
                        </Link>
                    </h1>
                    <div className="flex space-x-8">
                        <Link href="/schedule" className="text-xl group relative">
                            <span className="absolute inset-x-0 bottom-[-8px] 
                            h-1 bg-white scale-x-0 group-hover:scale-x-100 
                            group-hover:bg-[#BA1C21] transition-all duration-300">
                            </span>
                            SCHEDULE
                        </Link>
                        <Link href="/logout" className="text-xl group relative">
                            <span className="absolute inset-x-0 bottom-[-8px] 
                            h-1 bg-white scale-x-0 group-hover:scale-x-100 
                            group-hover:bg-[#BA1C21] transition-all duration-300">
                            </span>
                            LOGOUT
                        </Link>
                    </div>
                </div>
            </div>
        </div>
    );
}

