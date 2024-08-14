"use client";

import React, { useState } from "react";
import { Button, Label } from "flowbite-react";
import { HiMail } from "react-icons/hi";
import axios from "axios"; 
import { verifyEmail, verifyOtp } from "@/app/api/authAPI";

export default function SendVerificationCodePage() {
  const [email, setEmail] = useState("");
  const [sessionToken, setSessionToken] = useState("");
  const [verificationSent, setVerificationSent] = useState(false);
  const [errorMessage, setErrorMessage] = useState("");

  const handleSendVerificationCode = async () => {
    try {
     
      setErrorMessage("");

      const response = await axios.post("http://localhost:1323/verifyEmail", {
        email,
      });

      // If the request is successful, save the session token and update state
      setSessionToken(response.data.session);
      setVerificationSent(true);

      console.log("Verification code sent to:", email);
    } catch (error: any) {
   
      setErrorMessage(
        error.response?.data?.error || "Failed to send verification code."
      );
    }
  };

  return (
    <div className="container mx-auto py-8 flex flex-col items-center justify-center min-h-screen">
      <form className="w-full max-w-lg mx-auto bg-white p-8 rounded-md shadow-md">
        <div className="flex justify-center mb-6">
          <img src="/logo.png" alt="My Logo" width="150" height="150" />
        </div>
        <h1 className="font-bold mb-3 text-center">
          Welcome to Panchang API Gateway!
        </h1>

        <div className="mb-4">
          <Label
            className="block text-gray-700 text-sm font-bold mb-2"
            htmlFor="email"
          >
            Enter your email
          </Label>
          <div className="relative">
            <input
              className="w-full pl-10 pr-4 py-2 border-gray-300 rounded-md focus:outline-none focus:border-blue-500"
              id="email"
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              placeholder="Enter your email"
            />
            <HiMail className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" />
          </div>
        </div>

        {verificationSent && (
          <p className="text-green-500 mb-4">
            Verification code sent successfully!
          </p>
        )}

        {errorMessage && <p className="text-red-500 mb-4">{errorMessage}</p>}

        <Button
          onClick={handleSendVerificationCode}
          size="xs"
          className="w-full text-white bg-[#723B13] text-xl font-bold py-3 rounded-md transition duration-300"
        >
          Send Verification Code
        </Button>
      </form>
    </div>
  );
}
