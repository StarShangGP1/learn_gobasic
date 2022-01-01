package main

import (
	"learn_go/class11/fatrates"
	"testing"
)

var frk = &fatrates.FatRateRank{}

func TestCase1(t *testing.T) {
	{
		frk.InputRecord("wq", 0.38)
		frk.InputRecord("wq", 0.32)
		randWq, fatRateWq := frk.GetRank("wq")
		if randWq != 1 {
			t.Fatalf("wq 不是第一，%d", randWq)
		}
		if fatRateWq != 0.32 {
			t.Fatalf("wq 体脂不是0.32，%f", fatRateWq)
		}
	}

	frk.InputRecord("lj", 0.28)
	{
		randWq, fatRateWq := frk.GetRank("wq")
		if randWq != 2 {
			t.Fatalf("wq 不是第二，%d", randWq)
		}
		if fatRateWq != 0.32 {
			t.Fatalf("wq 体脂不是0.32，%f", fatRateWq)
		}
	}
	{

		randLj, fatRateLj := frk.GetRank("lj")
		if randLj != 1 {
			t.Fatalf("lj 不是第一，%d", randLj)
		}
		if fatRateLj != 0.28 {
			t.Fatalf("lj 体脂不是0.28，%f", fatRateLj)
		}
	}
}

func TestCase2(t *testing.T) {

	frk.InputRecord("wq", 0.38)
	frk.InputRecord("zw", 0.38)
	frk.InputRecord("lj", 0.28)
	{
		randLj, fatRateLj := frk.GetRank("lj")
		if randLj != 1 {
			t.Fatalf("lj 不是第一，%d", randLj)
		}
		if fatRateLj != 0.28 {
			t.Fatalf("lj 体脂不是0.28，%f", fatRateLj)
		}
	}
	{
		randWq, fatRateWq := frk.GetRank("wq")
		if randWq != 2 {
			t.Fatalf("wq 不是第二，%d", randWq)
		}
		if fatRateWq != 0.38 {
			t.Fatalf("wq 体脂不是0.32，%f", fatRateWq)
		}
	}
	{
		randZw, fatRateZw := frk.GetRank("wq")
		if randZw != 2 {
			t.Fatalf("zw 不是第二，%d", randZw)
		}
		if fatRateZw != 0.38 {
			t.Fatalf("zw 体脂不是0.32，%f", fatRateZw)
		}
	}
}

func TestCase3(t *testing.T) {

	frk.InputRecord("wq", 0.38)
	frk.InputRecord("lj", 0.28)
	frk.InputRecord("zw")
	{
		randLj, fatRateLj := frk.GetRank("lj")
		if randLj != 1 {
			t.Fatalf("lj 不是第一，%d", randLj)
		}
		if fatRateLj != 0.28 {
			t.Fatalf("lj 体脂不是0.28，%f", fatRateLj)
		}
	}
	{
		randWq, fatRateWq := frk.GetRank("wq")
		if randWq != 2 {
			t.Fatalf("wq 不是第二，%d", randWq)
		}
		if fatRateWq != 0.38 {
			t.Fatalf("wq 体脂不是0.32，%f", fatRateWq)
		}
	}
	{
		randZw, _ := frk.GetRank("zw")
		if randZw != 3 {
			t.Fatalf("zw 不是第三，%d", randZw)
		}
	}
}
