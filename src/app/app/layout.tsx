import type {Metadata} from "next";
import {Geist, Geist_Mono} from "next/font/google";
import "./globals.css";

const geistSans = Geist({
    variable: "--font-geist-sans",
    subsets: ["latin"],
});

const geistMono = Geist_Mono({
    variable: "--font-geist-mono",
    subsets: ["latin"],
});

export const metadata: Metadata = {
    title: "Passkey Study",
};

export default function RootLayout({
                                       children,
                                   }: Readonly<{
    children: React.ReactNode;
}>) {
    return (
        <html lang="en">
        <body>
        {/* nav var*/}
        <header>
            <nav className={"bg-green-500"}>
                <div className={"container mx-auto p-4"}>
                    <div className={"flex justify-between"}>
                        <div className={"flex"}>
                            <a href={"/"} className={"text-white font-bold"}>Passkey Study</a>
                        </div>
                        <div className={"flex"}>
                            <a href={"/"} className={"text-white"}>Home</a>
                            <a href={"/signup"} className={"text-white ml-4"}>Sign Up</a>
                            <a href={"/login"} className={"text-white ml-4"}>Login</a>
                        </div>
                    </div>
                </div>
            </nav>
        </header>
        {/* main content */}
        <main className={"container mx-auto p-4"}>{children}</main>
        {/* footer */}
        <footer className={"bg-indigo-950 mt-auto"}>
            <div className={"container mx-auto p-4"}>
                <div className={"flex justify-between text-white"}>
                    <div className={"flex"}>
                        <p>2025 Passkey Study</p>
                    </div>
                    <div className={"flex"}>
                        <a href={"/"} className={"text-white"}>Home</a>
                        <a href={"/about"} className={"text-white ml-4"}>About</a>
                    </div>
                </div>
            </div>
        </footer>
        </body>
        </html>
    );
}
