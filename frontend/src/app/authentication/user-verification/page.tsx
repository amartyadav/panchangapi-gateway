"use client";

import React, { useState } from "react";
import { Button, Label, Spinner } from "flowbite-react";
import { HiMail } from "react-icons/hi";
import axios from "axios";
import { useRouter } from "next/navigation";

export default function SendVerificationCodePage() {
  const [email, setEmail] = useState("");
  const [errorMessage, setErrorMessage] = useState("");
  const [successMessage, setSuccessMessage] = useState("");
  const [loading, setLoading] = useState(false); // Loading state
  const router = useRouter();

  const handleSendVerificationCode = async () => {
    setLoading(true); // Start loading when the button is clicked
    setErrorMessage("");
    setSuccessMessage("");

    try {
      const response = await axios.post(
        "http://localhost:1323/sendverificationemail",
        { email }
      );

      // Extract the session token from the response
      const sessionToken = response.data.session;

      // Save the session token in localStorage
      localStorage.setItem("sessionToken", sessionToken);

      // Set the success message
      setSuccessMessage("Verification code sent successfully!");

      console.log("Verification code sent to:", email);
    } catch (error: any) {
      setErrorMessage(
        error.response?.data?.error || "Failed to send verification code."
      );
    } finally {
      setLoading(false); // Stop loading once the request is complete
    }
  };

  const handleProceedToVerification = () => {
    // Redirect to the verify-email page after the user acknowledges the success message
    router.push("/authentication/verify-email");
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
              disabled={loading} // Disable input while loading
            />
            <HiMail className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400" />
          </div>
        </div>

        {errorMessage && <p className="text-red-500 mb-4">{errorMessage}</p>}
        {successMessage && (
          <div className="text-green-500 mb-4">
            <p>{successMessage}</p>
            <Button
              onClick={handleProceedToVerification}
              size="xs"
              className="mt-4 w-full text-white bg-[#723B13] text-xl font-bold py-3 rounded-md transition duration-300"
            >
              OK
            </Button>
          </div>
        )}

        {!successMessage && (
          <Button
            onClick={handleSendVerificationCode}
            size="xs"
            className="w-full text-white bg-[#723B13] text-xl font-bold py-3 rounded-md transition duration-300 flex items-center justify-center"
            disabled={loading} // Disable button while loading
          >
            {loading ? <Spinner size="sm" className="mr-2" /> : null}
            {loading ? "Sending..." : "Send Verification Code"}
          </Button>
        )}
      </form>
    </div>
  );
}
