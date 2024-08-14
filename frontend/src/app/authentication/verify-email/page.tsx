"use client";

import React, { useState } from "react";
import { Button, Label } from "flowbite-react";
import { HiLockClosed } from "react-icons/hi";

export default function VerifyEmailPage() {
  const [verificationCode, setVerificationCode] = useState("");

  const handleVerifyCode = () => {
    console.log("Verification code entered:", verificationCode);
    // You can add logic to verify the code here
  };

  return (
    <div className="container mx-auto py-8 flex flex-col items-center justify-center min-h-screen">
      <form className="w-full max-w-lg mx-auto bg-white p-8 rounded-md shadow-md">
        <h1 className="font-bold mb-3 text-center text-xl">
          Verify Your Email
        </h1>
        <p className="text-center text-gray-600 mb-6">
          Please enter the verification code that was sent to your email.
        </p>

        <div className="mb-4">
          <Label
            className="block text-gray-700 text-sm font-bold mb-2"
            htmlFor="verificationCode"
          >
            Verification Code
          </Label>
          <div className="relative">
            <input
              className="w-full pl-10 pr-4 py-2 border-gray-300 rounded-md focus:outline-none focus:border-blue-500"
              id="verificationCode"
              type="text"
              value={verificationCode}
              onChange={(e) => setVerificationCode(e.target.value)}
              placeholder="Enter verification code"
            />
            <HiLockClosed className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" />
          </div>
        </div>

        <Button
          onClick={handleVerifyCode}
          color="orange"
          size="xs"
          className="w-full text-white bg-[#723B13] text-xl font-bold py-3 rounded-md transition duration-300"
        >
          Verify Email
        </Button>
      </form>
    </div>
  );
}
