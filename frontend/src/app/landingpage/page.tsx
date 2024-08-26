"use client"

import React from "react";
import { Button } from "flowbite-react";
import Link from "next/link";

export default function LandingPage() {
  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      {/* Modal-like Container */}
      <div className="bg-white p-10 rounded-md shadow-lg w-full max-w-md">
        {/* Logo */}
        <div className="flex justify-center mb-6">
          <img src="/logo.png" alt="My Logo" width="150" height="150" />
        </div>

        {/* API Summary */}
        <h1 className="text-center font-bold text-2xl mb-3">
          Welcome to Panchang API Gateway!
        </h1>
        <p className="text-center text-gray-700 mb-6">
          Panchang API offers reliable and easy access to panchang data for
          your applications. Simplify your integration with our powerful tools.
        </p>

        {/* Buttons */}
        <div className="flex justify-center space-x-4">
          <Link href="/authentication/user-verification" passHref>
            <Button
              size="lg"
              className="text-white bg-[#723B13] font-bold py-3 px-6 rounded-md transition duration-300"
            >
              Sign Up
            </Button>
          </Link>
          <Link href="/authentication/user-login" passHref>
            <Button
              size="lg"
              className="text-white bg-[#723B13] font-bold py-3 px-6 rounded-md transition duration-300"
            >
              Login
            </Button>
          </Link>
        </div>
      </div>
    </div>
  );
}
