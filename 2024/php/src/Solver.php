<?php

declare(strict_types=1);

namespace Adventofcode24;

interface Solver
{
    public function __construct(string $input);

    public function solvePart1(): string;

    public function solvePart2(): string;
}
