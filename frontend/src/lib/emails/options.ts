import gmail from "@/assets/images/gmail.png";
import outlook from "@/assets/images/outlook.png";

export type EmailOptions = {
  title: "Gmail" | "Outlook";
  icon: string;
  disabled?: boolean;
}[];

export const EmailOptions: EmailOptions = [
  {
    title: "Gmail",
    icon: gmail,
  },
  {
    title: "Outlook",
    icon: outlook,
    disabled: true,
  },
];
