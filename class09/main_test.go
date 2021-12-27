package main

import "testing"

func TestCase1(t *testing.T) {
	{
		inputRecord("wq", 0.38)
		inputRecord("wq", 0.32)
		randWq, fatRateWq := getRank("wq")
		if randWq != 1 {
			t.Fatalf("wq 不是第一，%d", randWq)
		}
		if fatRateWq != 0.32 {
			t.Fatalf("wq 体脂不是0.32，%f", fatRateWq)
		}
	}

	inputRecord("lj", 0.28)
	{
		randWq, fatRateWq := getRank("wq")
		if randWq != 2 {
			t.Fatalf("wq 不是第二，%d", randWq)
		}
		if fatRateWq != 0.32 {
			t.Fatalf("wq 体脂不是0.32，%f", fatRateWq)
		}
	}
	{

		randLj, fatRateLj := getRank("lj")
		if randLj != 1 {
			t.Fatalf("lj 不是第一，%d", randLj)
		}
		if fatRateLj != 0.28 {
			t.Fatalf("lj 体脂不是0.28，%f", fatRateLj)
		}
	}
}

func TestCase2(t *testing.T) {

	inputRecord("wq", 0.38)
	inputRecord("zw", 0.38)
	inputRecord("lj", 0.28)
	{
		randLj, fatRateLj := getRank("lj")
		if randLj != 1 {
			t.Fatalf("lj 不是第一，%d", randLj)
		}
		if fatRateLj != 0.28 {
			t.Fatalf("lj 体脂不是0.28，%f", fatRateLj)
		}
	}
	{
		randWq, fatRateWq := getRank("wq")
		if randWq != 2 {
			t.Fatalf("wq 不是第二，%d", randWq)
		}
		if fatRateWq != 0.38 {
			t.Fatalf("wq 体脂不是0.32，%f", fatRateWq)
		}
	}
	{
		randZw, fatRateZw := getRank("wq")
		if randZw != 2 {
			t.Fatalf("zw 不是第二，%d", randZw)
		}
		if fatRateZw != 0.38 {
			t.Fatalf("zw 体脂不是0.32，%f", fatRateZw)
		}
	}
}

func TestCase3(t *testing.T) {

	inputRecord("wq", 0.38)
	inputRecord("lj", 0.28)
	inputRecord("zw")
	{
		randLj, fatRateLj := getRank("lj")
		if randLj != 1 {
			t.Fatalf("lj 不是第一，%d", randLj)
		}
		if fatRateLj != 0.28 {
			t.Fatalf("lj 体脂不是0.28，%f", fatRateLj)
		}
	}
	{
		randWq, fatRateWq := getRank("wq")
		if randWq != 2 {
			t.Fatalf("wq 不是第二，%d", randWq)
		}
		if fatRateWq != 0.38 {
			t.Fatalf("wq 体脂不是0.32，%f", fatRateWq)
		}
	}
	{
		randZw, _ := getRank("zw")
		if randZw != 3 {
			t.Fatalf("zw 不是第三，%d", randZw)
		}
	}
}
