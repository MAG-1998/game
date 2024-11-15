package utils

import (
	"math/rand"
	"time"
)

func GetRandomPhraseforheal() string {
	phrases := []string{"Грызем Семки", "Залпом выпиваем Пивко", "Кола!", "Завариваем Роллтон", "Отбираем у мимо проходящего курьера Суши", "Гумма номер 1"}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(phrases))
	return phrases[randomIndex]
}


func GetRandomPhrase() string {
	phrases := []string{"Че будем?", "Делай! Делай!", "Ехалооо!", "Снова в бой, гопник!", "Ботаник снова на ногах", "Выбери че надо"}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(phrases))
	return phrases[randomIndex]
}

func GetRandomPhraseforattack() string {
	phrases := []string{"Ебашим", "Пинаем", "Дубасим", "Кидаем на прогиб", "Правый хук", "Левый хук", "Аперкот", "Шапалак,", "Рыгаем на лицо"}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(phrases))
	return phrases[randomIndex]
}