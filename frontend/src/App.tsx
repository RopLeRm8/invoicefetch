import { useEffect, useRef, useState } from "react";
import "./App.css";
import { EmailOptions } from "./lib/config/emails/options";
import { cn } from "./utils/cn";
import auth from "@/assets/images/actions/auth.png";
import verify from "@/assets/images/actions/verify.png";
import fetch from "@/assets/images/actions/fetch.png";
import logout from "@/assets/images/actions/logout.png";
import { Auth, GetActiveEmail, Logout } from "@/../wailsjs/go/main/App";
import { EventsOn } from "@/../wailsjs/runtime/runtime";

const authenticate = async (provider: EmailOptions[number]["title"]) => {
  const { ok, url, error } = await Auth(provider);
  console.log(url, error);
};

function App() {
  const [emailOption, setEmailOption] = useState<
    EmailOptions[number]["title"] | undefined
  >(undefined);
  const [activeEmail, setActiveEmail] = useState<string | undefined>(undefined);

  const emailCoverRef = useRef<HTMLDivElement | null>(null);
  const mainDivRef = useRef<HTMLDivElement | null>(null);
  const emailOptionsRef = useRef<HTMLButtonElement[]>([]);

  const addEmailOptionRef = (el: HTMLButtonElement | null) => {
    if (el) emailOptionsRef.current.push(el);
  };

  const EmailLogout = async () => {
    const { error } = await Logout();
    if (error) console.log(error);
    setActiveEmail(undefined);
  };

  useEffect(() => {
    const emailEl = emailOptionsRef.current[0];
    const mainBoxEl = mainDivRef.current;
    if (!emailCoverRef.current || !emailEl || !mainBoxEl) return;
    const { left: mainBoxLeft } = mainBoxEl.getBoundingClientRect();
    const { left, width } = emailEl.getBoundingClientRect();

    const totalLeft = left - mainBoxLeft;
    emailCoverRef.current.style.left = `${totalLeft}px`;
    emailCoverRef.current.style.width = `${width}px`;
    setEmailOption("Gmail");

    EventsOn("auth:success", ({ email }: { email: string }) => {
      setActiveEmail(email);
    });
  }, []);

  useEffect(() => {
    const getEmail = async () => {
      const { email, error } = await GetActiveEmail();
      if (error) return;
      setActiveEmail(email);
    };
    getEmail();
  }, []);

  const selectEmail = (title: EmailOptions[number]["title"], ind: number) => {
    if (!emailCoverRef.current) return;
    const emailEl = emailOptionsRef.current[ind];
    const mainBoxEl = mainDivRef.current;
    if (!emailEl || !mainBoxEl) return;

    const { left: mainBoxLeft } = mainBoxEl.getBoundingClientRect();
    const { left, width } = emailEl.getBoundingClientRect();

    const totalLeft = left - mainBoxLeft;
    emailCoverRef.current.style.left = `${totalLeft}px`;
    emailCoverRef.current.style.width = `${width}px`;

    setEmailOption(title);
  };

  return (
    <main>
      <div className="pt-[5%] w-11/12 mx-auto flex h-screen relative">
        <div className="flex flex-col gap-1 flex-1 items-start">
          <div className="flex flex-col gap-1 text-start">
            <span className="text-4xl font-bold">Invoice Fetcher</span>
            <span className="opacity-80 font-light">
              Your ultimate email fetcher
            </span>
          </div>
          <div
            className="flex items-center h-10 relative w-fit mt-2"
            ref={mainDivRef}
          >
            <div
              className="absolute w-full h-full pointer-events-none bg-white duration-200 rounded-tl-2xl rounded-br-2xl"
              ref={emailCoverRef}
            />
            {EmailOptions.map((opt, ind) => {
              const isActive = emailOption === opt.title;
              const isFirst = ind === 0;
              const isLast = ind === EmailOptions.length - 1;
              return (
                <button
                  className={cn(
                    "flex items-center gap-2 text-white h-10 bg-black/60 px-6 duration-200 disabled:opacity-60",
                    isActive ? "text-black" : "text-white",
                    isFirst ? "rounded-l-2xl" : "",
                    isLast ? "rounded-r-2xl" : ""
                  )}
                  key={opt.title}
                  onClick={() => selectEmail(opt.title, ind)}
                  ref={(el) => addEmailOptionRef(el)}
                  disabled={opt.disabled}
                >
                  <img src={opt.icon} alt="" className="w-6 z-10" />
                  <span className="z-10">{opt.title}</span>
                </button>
              );
            })}
          </div>
          <div className="flex flex-col items-start mt-12 gap-2 w-48">
            <span className="text-2xl font-semibold mb-2">
              What to do next?
            </span>
            <button
              className="py-3 px-5 w-full bg-background hover:bg-opacity-40 duration-200 rounded-md flex items-center gap-6 disabled:opacity-70"
              onClick={() => authenticate("Gmail")}
              disabled={!!activeEmail}
            >
              <img src={auth} className="w-6" />
              Authenticate
            </button>
            <button className="py-3 px-5 w-full bg-background hover:bg-opacity-40 duration-200 rounded-md flex items-center gap-6">
              <img src={verify} className="w-6" />
              Verify Email
            </button>
            <button className="py-3 px-5 w-full bg-background hover:bg-opacity-40 duration-200 rounded-md flex items-center gap-6">
              <img src={fetch} className="w-6" />
              Fetch Emails
            </button>
            <button
              className="py-3 px-5 w-full bg-background hover:bg-opacity-40 duration-200 rounded-md flex items-center gap-6 disabled:opacity-70"
              onClick={EmailLogout}
              disabled={!activeEmail}
            >
              <img src={logout} className="w-6" />
              Logout
            </button>
          </div>
        </div>
        <div className="flex-1 ">
          <div className="w-full bg-background/60 h-[90%] shadow-sm shadow-black rounded-md">
            <div className="w-full h-[10%] bg-background flex justify-center items-center shadow-sm shadow-black">
              {activeEmail ? `Emails for ${activeEmail}` : "Not logged in"}
            </div>
            <div className="flex items-center h-[90%] justify-center">
              Output will be shown here
            </div>
          </div>
        </div>
      </div>
      <div className="absolute bottom-1 left-[1%] w-[98%] mx-auto flex justify-between">
        <span>In active development</span>
        <span>V1.0</span>
      </div>
    </main>
  );
}

export default App;
