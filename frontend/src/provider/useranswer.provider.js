import ApiHandler from "./api.handler";

const USER_ANSWER = "user-answer";

// 🚀 Submit user answer
export const submitUserAnswer = async (payload) => {
  const res = await ApiHandler.post(`/${USER_ANSWER}/`, payload);
  return res.data;
};
