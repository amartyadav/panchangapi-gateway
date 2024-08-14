import axios, { AxiosResponse } from "axios";

const BACKEND_URL = "http://localhost:1323";

export const verifyEmail = async (email: string): Promise<string> => {
  try {
    const response: AxiosResponse<{ session: string }> = await axios.post(
      `${BACKEND_URL}/sendverificationemail`,
      { email }
    );
    return response.data.session;
  } catch (error: any) {
    if (error.response) {
      const errorMessage =
        error.response.data.error || "Unknown error occurred";
      console.error("Error verifying email:", errorMessage);
      throw new Error(errorMessage);
    } else {
      console.error("Error verifying email:", error.message);
      throw new Error("Network error");
    }
  }
};

export const verifyOtp = async (
  sessionToken: string,
  otp: string
): Promise<void> => {
  try {
    await axios.post(`${BACKEND_URL}/verifyotp`, {
      sessionToken,
      otp,
    });
  } catch (error: any) {
    if (error.response) {
      const errorMessage =
        error.response.data.error || "Unknown error occurred";
      console.error("Error verifying OTP:", errorMessage);
      throw new Error(errorMessage);
    } else {
      console.error("Error verifying OTP:", error.message);
      throw new Error("Network error");
    }
  }
};
