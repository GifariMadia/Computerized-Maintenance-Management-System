import { NextResponse } from "next/server";

export async function GET() {
  return NextResponse.json(
    {
      status: "ok",
      service: "cmms-frontend",
      timezone: process.env.TZ ?? "Asia/Jakarta",
    },
    { status: 200 }
  );
}
